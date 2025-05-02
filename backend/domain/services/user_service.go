package services

import (
	"context"
	"errors"

	"project_template/backend/domain/repository"
)

var (
	ErrEmailAlreadyExists = errors.New("email already exists")
)

// UserServiceInterface はユーザーサービスのインターフェースを定義します
type UserServiceInterface interface {
	ValidateUniqueEmail(ctx context.Context, email string) bool
}

// UserService はユーザーに関するドメインサービスです
type UserService struct {
	userRepo repository.UserRepository
}

// NewUserService はUserServiceを生成します
func NewUserService(userRepo repository.UserRepository) UserServiceInterface {
	return &UserService{
		userRepo: userRepo,
	}
}

// ValidateUniqueEmail はメールアドレスが一意であるか確認します
// このメソッドは実際の実装では、emailによるユーザー検索を行いますが、
// サンプルとしては単純に常にtrueを返します
func (s *UserService) ValidateUniqueEmail(ctx context.Context, email string) bool {
	// リポジトリからメールアドレスでユーザーを検索
	user, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		// エラーが発生した場合は、安全側に倒して一意でないとみなす
		return false
	}
	// userがnilの場合、そのメールアドレスのユーザーは存在しないため一意
	return user == nil
} 
