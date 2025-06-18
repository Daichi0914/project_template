package repository

import (
	"context"

	"project_template/backend/domain/entity"
)

// UserRepository はユーザーのリポジトリインターフェースです
type UserRepository interface {
	FindByID(ctx context.Context, id string) (*entity.User, error)
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	FindAll(ctx context.Context) ([]*entity.User, error)
	Create(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id string) error
} 
