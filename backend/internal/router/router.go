package router

import (
	"github.com/Daichi0914/project_template/handler"
	"net/http"
)

func New() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", handler.HealthCheck)
	return mux
}
