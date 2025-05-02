package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"project_template/backend/domain/services"
	"project_template/backend/usecase/dto"
	"project_template/backend/usecase/interactor"
)

// MockUserInteractor はUserInteractorのモック実装です
type MockUserInteractor struct {
	mock.Mock
}

func (m *MockUserInteractor) GetUser(ctx context.Context, input *dto.GetUserInput) (*dto.UserOutput, error) {
	args := m.Called(ctx, input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.UserOutput), args.Error(1)
}

func (m *MockUserInteractor) CreateUser(ctx context.Context, input *dto.CreateUserInput) (*dto.UserOutput, error) {
	args := m.Called(ctx, input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.UserOutput), args.Error(1)
}

func TestGetUser_Success(t *testing.T) {
	// モックの設定
	mockInteractor := new(MockUserInteractor)
	createdAt, _ := time.Parse(time.RFC3339, "2025-05-03T04:46:51Z")
	updatedAt, _ := time.Parse(time.RFC3339, "2025-05-03T04:46:51Z")
	
	expectedOutput := &dto.UserOutput{
		ID:        "1",
		Name:      "サンプルユーザー",
		Email:     "sample@example.com",
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
	
	mockInteractor.On("GetUser", mock.Anything, &dto.GetUserInput{ID: "1"}).Return(expectedOutput, nil)
	
	// ハンドラーの作成
	handler := NewUserHandler(mockInteractor)
	
	// リクエストの作成
	req, _ := http.NewRequest("GET", "/api/v1/users/1", nil)
	// gorilla/muxのルーティングパラメータを設定
	vars := map[string]string{
		"id": "1",
	}
	req = mux.SetURLVars(req, vars)
	
	// レスポンスレコーダーの作成
	rr := httptest.NewRecorder()
	
	// ハンドラー関数の実行
	handler.GetUser(rr, req)
	
	// レスポンスの検証
	assert.Equal(t, http.StatusOK, rr.Code)
	
	var response dto.UserOutput
	json.Unmarshal(rr.Body.Bytes(), &response)
	
	assert.Equal(t, expectedOutput.ID, response.ID)
	assert.Equal(t, expectedOutput.Name, response.Name)
	assert.Equal(t, expectedOutput.Email, response.Email)
	
	// モックの検証
	mockInteractor.AssertExpectations(t)
}

func TestGetUser_NotFound(t *testing.T) {
	// モックの設定
	mockInteractor := new(MockUserInteractor)
	mockInteractor.On("GetUser", mock.Anything, &dto.GetUserInput{ID: "999"}).Return(nil, interactor.ErrUserNotFound)
	
	// ハンドラーの作成
	handler := NewUserHandler(mockInteractor)
	
	// リクエストの作成
	req, _ := http.NewRequest("GET", "/api/v1/users/999", nil)
	vars := map[string]string{
		"id": "999",
	}
	req = mux.SetURLVars(req, vars)
	
	// レスポンスレコーダーの作成
	rr := httptest.NewRecorder()
	
	// ハンドラー関数の実行
	handler.GetUser(rr, req)
	
	// レスポンスの検証
	assert.Equal(t, http.StatusNotFound, rr.Code)
	
	var response map[string]string
	json.Unmarshal(rr.Body.Bytes(), &response)
	
	assert.Equal(t, "user not found", response["error"])
	
	// モックの検証
	mockInteractor.AssertExpectations(t)
}

func TestCreateUser_Success(t *testing.T) {
	// モックの設定
	mockInteractor := new(MockUserInteractor)
	now := time.Now().UTC()
	
	input := &dto.CreateUserInput{
		Name:  "新規ユーザー",
		Email: "new@example.com",
	}
	
	expectedOutput := &dto.UserOutput{
		ID:        "123",
		Name:      "新規ユーザー",
		Email:     "new@example.com",
		CreatedAt: now,
		UpdatedAt: now,
	}
	
	mockInteractor.On("CreateUser", mock.Anything, mock.MatchedBy(func(i *dto.CreateUserInput) bool {
		return i.Name == input.Name && i.Email == input.Email
	})).Return(expectedOutput, nil)
	
	// ハンドラーの作成
	handler := NewUserHandler(mockInteractor)
	
	// リクエストの作成
	reqBody, _ := json.Marshal(input)
	req, _ := http.NewRequest("POST", "/api/v1/users", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	
	// レスポンスレコーダーの作成
	rr := httptest.NewRecorder()
	
	// ハンドラー関数の実行
	handler.CreateUser(rr, req)
	
	// レスポンスの検証
	assert.Equal(t, http.StatusCreated, rr.Code)
	
	var response dto.UserOutput
	json.Unmarshal(rr.Body.Bytes(), &response)
	
	assert.Equal(t, expectedOutput.ID, response.ID)
	assert.Equal(t, expectedOutput.Name, response.Name)
	assert.Equal(t, expectedOutput.Email, response.Email)
	
	// モックの検証
	mockInteractor.AssertExpectations(t)
}

func TestCreateUser_InvalidRequest(t *testing.T) {
	// モックの設定
	mockInteractor := new(MockUserInteractor)
	
	// ハンドラーの作成
	handler := NewUserHandler(mockInteractor)
	
	// 不正なJSONでリクエストを作成
	reqBody := []byte(`{"name": "テスト"`) // 閉じブラケットがない不正なJSON
	req, _ := http.NewRequest("POST", "/api/v1/users", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	
	// レスポンスレコーダーの作成
	rr := httptest.NewRecorder()
	
	// ハンドラー関数の実行
	handler.CreateUser(rr, req)
	
	// レスポンスの検証
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	
	var response map[string]string
	json.Unmarshal(rr.Body.Bytes(), &response)
	
	assert.Equal(t, "invalid request body", response["error"])
}

func TestCreateUser_EmailAlreadyExists(t *testing.T) {
	// モックの設定
	mockInteractor := new(MockUserInteractor)
	
	input := &dto.CreateUserInput{
		Name:  "重複ユーザー",
		Email: "duplicate@example.com",
	}
	
	// メールアドレス重複エラーを返すように設定
	mockInteractor.On("CreateUser", mock.Anything, mock.MatchedBy(func(i *dto.CreateUserInput) bool {
		return i.Name == input.Name && i.Email == input.Email
	})).Return(nil, services.ErrEmailAlreadyExists)
	
	// ハンドラーの作成
	handler := NewUserHandler(mockInteractor)
	
	// リクエストの作成
	reqBody, _ := json.Marshal(input)
	req, _ := http.NewRequest("POST", "/api/v1/users", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	
	// レスポンスレコーダーの作成
	rr := httptest.NewRecorder()
	
	// ハンドラー関数の実行
	handler.CreateUser(rr, req)
	
	// レスポンスの検証
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	
	var response map[string]string
	json.Unmarshal(rr.Body.Bytes(), &response)
	
	assert.Equal(t, "email already exists", response["error"])
	
	// モックの検証
	mockInteractor.AssertExpectations(t)
} 