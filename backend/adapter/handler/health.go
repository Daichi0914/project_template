package handler

import (
	"fmt"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, "ok")
}
