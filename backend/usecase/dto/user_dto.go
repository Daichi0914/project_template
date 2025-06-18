package dto

import (
	"time"

	"project_template/backend/domain/entity"
)

// UserInput は新規ユーザー作成のための入力データです
type CreateUserInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// GetUserInput はユーザー取得のための入力データです
type GetUserInput struct {
	ID string `json:"id"`
}

// UserOutput はユーザー情報の出力データです
type UserOutput struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewUserOutput はエンティティからDTOへの変換を行います
func NewUserOutput(user *entity.User) *UserOutput {
	return &UserOutput{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

// UsersOutput はユーザー一覧の出力データです
type UsersOutput struct {
	Users []*UserOutput `json:"users"`
}

// NewUsersOutput はエンティティのスライスからDTOへの変換を行います
func NewUsersOutput(users []*entity.User) *UsersOutput {
	userOutputs := make([]*UserOutput, len(users))
	for i, user := range users {
		userOutputs[i] = NewUserOutput(user)
	}
	return &UsersOutput{
		Users: userOutputs,
	}
}
