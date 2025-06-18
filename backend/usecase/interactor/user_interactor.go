package interactor

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"project_template/backend/domain/entity"
	"project_template/backend/domain/repository"
	"project_template/backend/domain/services"
	"project_template/backend/usecase/dto"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

// UserInteractor はユーザーに関するユースケースを実装します
type UserInteractor struct {
	userRepo    repository.UserRepository
	userService services.UserServiceInterface
}

// NewUserInteractor はUserInteractorを生成します
func NewUserInteractor(
	userRepo repository.UserRepository,
	userService services.UserServiceInterface,
) *UserInteractor {
	return &UserInteractor{
		userRepo:    userRepo,
		userService: userService,
	}
}

// GetUser はユーザー情報を取得します
func (i *UserInteractor) GetUser(ctx context.Context, input *dto.GetUserInput) (*dto.UserOutput, error) {
	// リポジトリからユーザーを取得
	user, err := i.userRepo.FindByID(ctx, input.ID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}

	// ドメインオブジェクトをDTOに変換して返却
	return dto.NewUserOutput(user), nil
}

// GetUsers はすべてのユーザー情報を取得します
func (i *UserInteractor) GetUsers(ctx context.Context) (*dto.UsersOutput, error) {
	// リポジトリからすべてのユーザーを取得
	users, err := i.userRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	// ドメインオブジェクトをDTOに変換して返却
	return dto.NewUsersOutput(users), nil
}

// CreateUser は新規ユーザーを作成します
func (i *UserInteractor) CreateUser(ctx context.Context, input *dto.CreateUserInput) (*dto.UserOutput, error) {
	// メールアドレスの一意性を確認
	if !i.userService.ValidateUniqueEmail(ctx, input.Email) {
		return nil, services.ErrEmailAlreadyExists
	}

	// ユーザーエンティティを作成
	userID := uuid.New().String()
	user := entity.NewUser(userID, input.Name, input.Email)

	// リポジトリに保存
	err := i.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	// ドメインオブジェクトをDTOに変換して返却
	return dto.NewUserOutput(user), nil
}
