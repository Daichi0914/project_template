package services

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"project_template/backend/domain/entity"
)

// MockUserRepository はUserRepositoryのモック実装です
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) FindByID(ctx context.Context, id string) (*entity.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *MockUserRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	args := m.Called(ctx, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *MockUserRepository) Create(ctx context.Context, user *entity.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) Update(ctx context.Context, user *entity.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestValidateUniqueEmail_WhenUserDoesNotExist(t *testing.T) {
	// モックの設定
	mockRepo := new(MockUserRepository)
	mockRepo.On("FindByEmail", mock.Anything, "new@example.com").Return(nil, nil)
	
	// サービスの作成
	service := NewUserService(mockRepo)
	
	// テスト実行
	result := service.ValidateUniqueEmail(context.Background(), "new@example.com")
	
	// 検証 - ユーザーが存在しない場合はtrueを返すべき
	assert.True(t, result)
	
	// モックの検証
	mockRepo.AssertExpectations(t)
}

func TestValidateUniqueEmail_WhenUserExists(t *testing.T) {
	// モックの設定
	mockRepo := new(MockUserRepository)
	existingUser := &entity.User{
		ID:    "00000000-0000-0000-0000-000000000001",
		Email: "existing@example.com",
	}
	mockRepo.On("FindByEmail", mock.Anything, "existing@example.com").Return(existingUser, nil)
	
	// サービスの作成
	service := NewUserService(mockRepo)
	
	// テスト実行
	result := service.ValidateUniqueEmail(context.Background(), "existing@example.com")
	
	// 検証 - ユーザーが存在する場合はfalseを返すべき
	assert.False(t, result)
	
	// モックの検証
	mockRepo.AssertExpectations(t)
}

func TestValidateUniqueEmail_WhenErrorOccurs(t *testing.T) {
	// モックの設定
	mockRepo := new(MockUserRepository)
	mockRepo.On("FindByEmail", mock.Anything, "error@example.com").Return(nil, errors.New("database error"))
	
	// サービスの作成
	service := NewUserService(mockRepo)
	
	// テスト実行
	result := service.ValidateUniqueEmail(context.Background(), "error@example.com")
	
	// 検証 - エラーが発生した場合は安全側に倒してfalseを返すべき
	assert.False(t, result)
	
	// モックの検証
	mockRepo.AssertExpectations(t)
} 