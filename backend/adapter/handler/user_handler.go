package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"project_template/backend/adapter/middleware"
	"project_template/backend/domain/services"
	"project_template/backend/usecase/dto"
	"project_template/backend/usecase/interactor"
)

// UserInteractorInterface はユーザーインタラクターのインターフェースを定義します
type UserInteractorInterface interface {
	GetUser(ctx context.Context, input *dto.GetUserInput) (*dto.UserOutput, error)
	GetUsers(ctx context.Context) (*dto.UsersOutput, error)
	CreateUser(ctx context.Context, input *dto.CreateUserInput) (*dto.UserOutput, error)
}

// UserHandler はユーザー関連のHTTPリクエストを処理します
type UserHandler struct {
	userInteractor UserInteractorInterface
}

// NewUserHandler はUserHandlerを生成します
func NewUserHandler(userInteractor UserInteractorInterface) *UserHandler {
	return &UserHandler{
		userInteractor: userInteractor,
	}
}

// GetUser はユーザー情報を取得するハンドラーです
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	ctx := r.Context()
	input := &dto.GetUserInput{
		ID: userID,
	}

	output, err := h.userInteractor.GetUser(ctx, input)
	if err != nil {
		resp := middleware.NewJSONResponse(w)
		if err == interactor.ErrUserNotFound {
			resp.Encode(http.StatusNotFound, map[string]string{"error": "user not found"})
			return
		}
		resp.Encode(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
		return
	}

	// 名前のUTF-8検証と正規化
	if output != nil {
		output.Name = middleware.SanitizeString(output.Name)
	}
	
	resp := middleware.NewJSONResponse(w)
	resp.Encode(http.StatusOK, output)
}

// GetUsers はユーザー一覧を取得するハンドラーです
func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	output, err := h.userInteractor.GetUsers(ctx)
	if err != nil {
		resp := middleware.NewJSONResponse(w)
		resp.Encode(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
		return
	}

	// 各ユーザーの名前のUTF-8検証と正規化
	if output != nil && output.Users != nil {
		for _, user := range output.Users {
			if user != nil {
				user.Name = middleware.SanitizeString(user.Name)
			}
		}
	}
	
	resp := middleware.NewJSONResponse(w)
	resp.Encode(http.StatusOK, output)
}

// CreateUser は新規ユーザーを作成するハンドラーです
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateUserInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		resp := middleware.NewJSONResponse(w)
		resp.Encode(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}

	// 入力データの正規化
	input.Name = middleware.SanitizeString(input.Name)
	
	ctx := r.Context()
	output, err := h.userInteractor.CreateUser(ctx, &input)
	if err != nil {
		resp := middleware.NewJSONResponse(w)
		// エラーの種類によって適切なステータスコードを返す
		if err == services.ErrEmailAlreadyExists {
			resp.Encode(http.StatusBadRequest, map[string]string{"error": "email already exists"})
			return
		}
		resp.Encode(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
		return
	}

	resp := middleware.NewJSONResponse(w)
	resp.Encode(http.StatusCreated, output)
}
