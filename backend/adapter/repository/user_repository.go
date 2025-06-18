package repository

import (
	"context"
	"database/sql"
	"errors"

	"project_template/backend/domain/entity"
	domainRepo "project_template/backend/domain/repository"
)

var (
	ErrDBError = errors.New("database error")
)

// UserRepository はユーザーのリポジトリ実装です
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository はUserRepositoryを生成します
func NewUserRepository(db *sql.DB) domainRepo.UserRepository {
	return &UserRepository{
		db: db,
	}
}

// FindByID はIDによるユーザー検索を実装します
func (r *UserRepository) FindByID(ctx context.Context, id string) (*entity.User, error) {
	query := "SELECT id, name, email, created_at, updated_at FROM users WHERE id = ?"
	
	var user entity.User
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // ユーザーが見つからない場合
		}
		return nil, err
	}

	return &user, nil
}

// FindByEmail はメールアドレスによるユーザー検索を実装します
func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	// LOWER関数を使用して大文字小文字を区別せずに検索
	query := "SELECT id, name, email, created_at, updated_at FROM users WHERE LOWER(email) = LOWER(?)"
	
	var user entity.User
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // ユーザーが見つからない場合
		}
		return nil, err
	}

	return &user, nil
}

// FindAll はすべてのユーザーを取得します
func (r *UserRepository) FindAll(ctx context.Context) ([]*entity.User, error) {
	query := "SELECT id, name, email, created_at, updated_at FROM users ORDER BY created_at DESC"
	
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*entity.User
	for rows.Next() {
		var user entity.User
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// Create は新規ユーザーの保存を実装します
func (r *UserRepository) Create(ctx context.Context, user *entity.User) error {
	query := `INSERT INTO users (id, name, email, created_at, updated_at) 
			  VALUES (?, ?, ?, ?, ?)`
	
	_, err := r.db.ExecContext(
		ctx, 
		query,
		user.ID,
		user.Name,
		user.Email,
		user.CreatedAt,
		user.UpdatedAt,
	)
	
	if err != nil {
		return err
	}
	
	return nil
}

// Update はユーザー情報の更新を実装します
func (r *UserRepository) Update(ctx context.Context, user *entity.User) error {
	query := `UPDATE users 
			  SET name = ?, email = ?, updated_at = ? 
			  WHERE id = ?`
	
	result, err := r.db.ExecContext(
		ctx, 
		query,
		user.Name,
		user.Email,
		user.UpdatedAt,
		user.ID,
	)
	
	if err != nil {
		return err
	}
	
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	
	if rowsAffected == 0 {
		return errors.New("user not found")
	}
	
	return nil
}

// Delete はユーザーの削除を実装します
func (r *UserRepository) Delete(ctx context.Context, id string) error {
	query := "DELETE FROM users WHERE id = ?"
	
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	
	if rowsAffected == 0 {
		return errors.New("user not found")
	}
	
	return nil
}
