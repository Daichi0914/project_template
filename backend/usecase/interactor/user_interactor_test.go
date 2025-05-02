package interactor

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"project_template/backend/domain/entity"
	"project_template/backend/domain/services"
	"project_template/backend/usecase/dto"
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

// MockUserService はUserServiceのモック実装です
type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) ValidateUniqueEmail(ctx context.Context, email string) bool {
	args := m.Called(ctx, email)
	return args.Bool(0)
}

func TestGetUser_Success(t *testing.T) {
	// モックの設定
	mockRepo := new(MockUserRepository)
	
	createdAt := time.Now().UTC()
	updatedAt := time.Now().UTC()
	
	user := &entity.User{
		ID:        "1",
		Name:      "サンプルユーザー",
		Email:     "sample@example.com",
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
	
	mockRepo.On("FindByID", mock.Anything, "1").Return(user, nil)
	
	// サービスのモック
	mockService := new(MockUserService)
	
	// インタラクターの作成
	interactor := NewUserInteractor(mockRepo, mockService)
	
	// テスト実行
	result, err := interactor.GetUser(context.Background(), &dto.GetUserInput{ID: "1"})
	
	// 検証
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, user.ID, result.ID)
	assert.Equal(t, user.Name, result.Name)
	assert.Equal(t, user.Email, result.Email)
	assert.Equal(t, user.CreatedAt, result.CreatedAt)
	assert.Equal(t, user.UpdatedAt, result.UpdatedAt)
	
	// モックの検証
	mockRepo.AssertExpectations(t)
}

func TestGetUser_NotFound(t *testing.T) {
	// モックの設定
	mockRepo := new(MockUserRepository)
	
	mockRepo.On("FindByID", mock.Anything, "999").Return(nil, nil)
	
	// サービスのモック
	mockService := new(MockUserService)
	
	// インタラクターの作成
	interactor := NewUserInteractor(mockRepo, mockService)
	
	// テスト実行
	result, err := interactor.GetUser(context.Background(), &dto.GetUserInput{ID: "999"})
	
	// 検証
	assert.Error(t, err)
	assert.Equal(t, ErrUserNotFound, err)
	assert.Nil(t, result)
	
	// モックの検証
	mockRepo.AssertExpectations(t)
}

func TestCreateUser_Success(t *testing.T) {
	// モックの設定
	mockRepo := new(MockUserRepository)
	
	input := &dto.CreateUserInput{
		Name:  "新規ユーザー",
		Email: "new@example.com",
	}
	
	mockRepo.On("Create", mock.Anything, mock.AnythingOfType("*entity.User")).Return(nil)
	
	// サービスのモック - ValidateUniqueEmailがtrueを返すように設定
	mockService := new(MockUserService)
	mockService.On("ValidateUniqueEmail", mock.Anything, "new@example.com").Return(true)
	
	// インタラクターの作成
	interactor := NewUserInteractor(mockRepo, mockService)
	
	// テスト実行
	result, err := interactor.CreateUser(context.Background(), input)
	
	// 検証
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotEmpty(t, result.ID)
	assert.Equal(t, input.Name, result.Name)
	assert.Equal(t, input.Email, result.Email)
	assert.False(t, result.CreatedAt.IsZero())
	assert.False(t, result.UpdatedAt.IsZero())
	
	// モックの検証
	mockRepo.AssertExpectations(t)
	mockService.AssertExpectations(t)
}

func TestCreateUser_EmailAlreadyExists(t *testing.T) {
	// モックの設定
	mockRepo := new(MockUserRepository)
	
	input := &dto.CreateUserInput{
		Name:  "新規ユーザー",
		Email: "existing@example.com",
	}
	
	// サービスのモック - ValidateUniqueEmailがfalseを返すように設定
	mockService := new(MockUserService)
	mockService.On("ValidateUniqueEmail", mock.Anything, "existing@example.com").Return(false)
	
	// インタラクターの作成
	interactor := NewUserInteractor(mockRepo, mockService)
	
	// テスト実行
	result, err := interactor.CreateUser(context.Background(), input)
	
	// 検証
	assert.Error(t, err)
	assert.Equal(t, services.ErrEmailAlreadyExists, err)
	assert.Nil(t, result)
	
	// モックの検証
	mockService.AssertExpectations(t)
} 