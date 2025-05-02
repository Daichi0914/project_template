package repository

import (
	"context"
	"database/sql"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"project_template/backend/domain/entity"
)

func setupMockDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock, *UserRepository) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock DB: %v", err)
	}

	repo := &UserRepository{db: db}
	return db, mock, repo
}

func TestFindByID_Success(t *testing.T) {
	db, mock, repo := setupMockDB(t)
	defer db.Close()

	// テストデータ
	userID := "00000000-0000-0000-0000-000000000001"
	createdAt := time.Date(2025, 5, 3, 4, 46, 51, 0, time.UTC)
	updatedAt := time.Date(2025, 5, 3, 4, 46, 51, 0, time.UTC)

	// SQLクエリとその結果をモック
	rows := sqlmock.NewRows([]string{"id", "name", "email", "created_at", "updated_at"}).
		AddRow(userID, "サンプルユーザー", "sample@example.com", createdAt, updatedAt)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, email, created_at, updated_at FROM users WHERE id = ?")).
		WithArgs(userID).
		WillReturnRows(rows)

	// テスト実行
	user, err := repo.FindByID(context.Background(), userID)

	// 検証
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, userID, user.ID)
	assert.Equal(t, "サンプルユーザー", user.Name)
	assert.Equal(t, "sample@example.com", user.Email)
	assert.Equal(t, createdAt, user.CreatedAt)
	assert.Equal(t, updatedAt, user.UpdatedAt)

	// すべてのモックが呼び出されたことを確認
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestFindByID_NotFound(t *testing.T) {
	db, mock, repo := setupMockDB(t)
	defer db.Close()

	// 存在しないIDをクエリ
	userID := "non-existent-id"

	// SQLクエリとその結果をモック - 行がない結果を返す
	rows := sqlmock.NewRows([]string{"id", "name", "email", "created_at", "updated_at"})

	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, email, created_at, updated_at FROM users WHERE id = ?")).
		WithArgs(userID).
		WillReturnRows(rows)

	// テスト実行
	user, err := repo.FindByID(context.Background(), userID)

	// 検証 - エラーはないが、userはnilであるべき
	assert.NoError(t, err)
	assert.Nil(t, user)

	// すべてのモックが呼び出されたことを確認
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestFindByEmail_Success(t *testing.T) {
	db, mock, repo := setupMockDB(t)
	defer db.Close()

	// テストデータ
	email := "sample@example.com"
	userID := "00000000-0000-0000-0000-000000000001"
	createdAt := time.Date(2025, 5, 3, 4, 46, 51, 0, time.UTC)
	updatedAt := time.Date(2025, 5, 3, 4, 46, 51, 0, time.UTC)

	// SQLクエリとその結果をモック
	rows := sqlmock.NewRows([]string{"id", "name", "email", "created_at", "updated_at"}).
		AddRow(userID, "サンプルユーザー", email, createdAt, updatedAt)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, email, created_at, updated_at FROM users WHERE LOWER(email) = LOWER(?)")).
		WithArgs(email).
		WillReturnRows(rows)

	// テスト実行
	user, err := repo.FindByEmail(context.Background(), email)

	// 検証
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, userID, user.ID)
	assert.Equal(t, "サンプルユーザー", user.Name)
	assert.Equal(t, email, user.Email)
	assert.Equal(t, createdAt, user.CreatedAt)
	assert.Equal(t, updatedAt, user.UpdatedAt)

	// すべてのモックが呼び出されたことを確認
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestFindByEmail_CaseInsensitive(t *testing.T) {
	db, mock, repo := setupMockDB(t)
	defer db.Close()

	// テストデータ - 大文字含むメールアドレス
	searchEmail := "Sample@EXAMPLE.com"
	storedEmail := "sample@example.com"
	userID := "00000000-0000-0000-0000-000000000001"
	createdAt := time.Date(2025, 5, 3, 4, 46, 51, 0, time.UTC)
	updatedAt := time.Date(2025, 5, 3, 4, 46, 51, 0, time.UTC)

	// SQLクエリとその結果をモック
	rows := sqlmock.NewRows([]string{"id", "name", "email", "created_at", "updated_at"}).
		AddRow(userID, "サンプルユーザー", storedEmail, createdAt, updatedAt)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, email, created_at, updated_at FROM users WHERE LOWER(email) = LOWER(?)")).
		WithArgs(searchEmail).
		WillReturnRows(rows)

	// テスト実行 - 大文字小文字が異なるメールアドレスで検索
	user, err := repo.FindByEmail(context.Background(), searchEmail)

	// 検証 - 大文字小文字が異なってもユーザーが見つかるべき
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, userID, user.ID)
	assert.Equal(t, storedEmail, user.Email) // 返されるメールアドレスは元の形式

	// すべてのモックが呼び出されたことを確認
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestFindByEmail_NotFound(t *testing.T) {
	db, mock, repo := setupMockDB(t)
	defer db.Close()

	// 存在しないメールアドレスをクエリ
	email := "notfound@example.com"

	// SQLクエリとその結果をモック - 行がない結果を返す
	rows := sqlmock.NewRows([]string{"id", "name", "email", "created_at", "updated_at"})

	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, email, created_at, updated_at FROM users WHERE LOWER(email) = LOWER(?)")).
		WithArgs(email).
		WillReturnRows(rows)

	// テスト実行
	user, err := repo.FindByEmail(context.Background(), email)

	// 検証 - エラーはないが、userはnilであるべき
	assert.NoError(t, err)
	assert.Nil(t, user)

	// すべてのモックが呼び出されたことを確認
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCreate_Success(t *testing.T) {
	db, mock, repo := setupMockDB(t)
	defer db.Close()

	// テストデータ
	now := time.Now().UTC()
	user := &entity.User{
		ID:        "00000000-0000-0000-0000-000000000001",
		Name:      "新規ユーザー",
		Email:     "new@example.com",
		CreatedAt: now,
		UpdatedAt: now,
	}

	// SQLクエリのモック
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO users (id, name, email, created_at, updated_at) VALUES (?, ?, ?, ?, ?)")).
		WithArgs(user.ID, user.Name, user.Email, user.CreatedAt, user.UpdatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// テスト実行
	err := repo.Create(context.Background(), user)

	// 検証
	assert.NoError(t, err)

	// すべてのモックが呼び出されたことを確認
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdate_Success(t *testing.T) {
	db, mock, repo := setupMockDB(t)
	defer db.Close()

	// テストデータ
	now := time.Now().UTC()
	user := &entity.User{
		ID:        "00000000-0000-0000-0000-000000000001",
		Name:      "更新ユーザー",
		Email:     "updated@example.com",
		UpdatedAt: now,
	}

	// SQLクエリのモック
	mock.ExpectExec(regexp.QuoteMeta("UPDATE users SET name = ?, email = ?, updated_at = ? WHERE id = ?")).
		WithArgs(user.Name, user.Email, user.UpdatedAt, user.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	// テスト実行
	err := repo.Update(context.Background(), user)

	// 検証
	assert.NoError(t, err)

	// すべてのモックが呼び出されたことを確認
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdate_NotFound(t *testing.T) {
	db, mock, repo := setupMockDB(t)
	defer db.Close()

	// テストデータ（存在しないID）
	now := time.Now().UTC()
	user := &entity.User{
		ID:        "non-existent-id",
		Name:      "存在しないユーザー",
		Email:     "notfound@example.com",
		UpdatedAt: now,
	}

	// SQLクエリのモック - 影響を受けた行数が0
	mock.ExpectExec(regexp.QuoteMeta("UPDATE users SET name = ?, email = ?, updated_at = ? WHERE id = ?")).
		WithArgs(user.Name, user.Email, user.UpdatedAt, user.ID).
		WillReturnResult(sqlmock.NewResult(0, 0))

	// テスト実行
	err := repo.Update(context.Background(), user)

	// 検証
	assert.Error(t, err)
	assert.Equal(t, "user not found", err.Error())

	// すべてのモックが呼び出されたことを確認
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDelete_Success(t *testing.T) {
	db, mock, repo := setupMockDB(t)
	defer db.Close()

	// テストデータ
	userID := "00000000-0000-0000-0000-000000000001"

	// SQLクエリのモック
	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM users WHERE id = ?")).
		WithArgs(userID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	// テスト実行
	err := repo.Delete(context.Background(), userID)

	// 検証
	assert.NoError(t, err)

	// すべてのモックが呼び出されたことを確認
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDelete_NotFound(t *testing.T) {
	db, mock, repo := setupMockDB(t)
	defer db.Close()

	// テストデータ（存在しないID）
	userID := "non-existent-id"

	// SQLクエリのモック - 影響を受けた行数が0
	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM users WHERE id = ?")).
		WithArgs(userID).
		WillReturnResult(sqlmock.NewResult(0, 0))

	// テスト実行
	err := repo.Delete(context.Background(), userID)

	// 検証
	assert.Error(t, err)
	assert.Equal(t, "user not found", err.Error())

	// すべてのモックが呼び出されたことを確認
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
} 