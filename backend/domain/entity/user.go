package entity

import (
	"time"
)

// User はユーザーを表すエンティティです
type User struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewUser はユーザーエンティティを生成します
func NewUser(id, name, email string) *User {
	now := time.Now()
	return &User{
		ID:        id,
		Name:      name,
		Email:     email,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// ChangeName はユーザー名を変更します
func (u *User) ChangeName(name string) {
	u.Name = name
	u.UpdatedAt = time.Now()
}

// ChangeEmail はメールアドレスを変更します
func (u *User) ChangeEmail(email string) {
	u.Email = email
	u.UpdatedAt = time.Now()
} 
