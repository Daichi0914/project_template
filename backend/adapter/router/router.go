package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"project_template/backend/adapter/handler"
)

// Router はアプリケーションのルーターを設定します
type Router struct {
	userHandler *handler.UserHandler
}

// NewRouter はRouterを生成します
func NewRouter(userHandler *handler.UserHandler) *Router {
	return &Router{
		userHandler: userHandler,
	}
}

// Setup はルーターの設定を行います
func (r *Router) Setup() *mux.Router {
	router := mux.NewRouter()

	// APIのバージョンプレフィックス
	api := router.PathPrefix("/api/v1").Subrouter()

	// ユーザー関連のエンドポイント
	api.HandleFunc("/users", r.userHandler.CreateUser).Methods(http.MethodPost)
	api.HandleFunc("/users/{id}", r.userHandler.GetUser).Methods(http.MethodGet)

	// ヘルスチェック
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods(http.MethodGet)

	return router
}
