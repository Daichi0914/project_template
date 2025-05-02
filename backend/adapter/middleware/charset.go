package middleware

import (
	"encoding/json"
	"net/http"
	"unicode/utf8"
)

// JSONResponse はJSON形式のレスポンスを表します
type JSONResponse struct {
	writer http.ResponseWriter
}

// NewJSONResponse は新しいJSONResponseを生成します
func NewJSONResponse(w http.ResponseWriter) *JSONResponse {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return &JSONResponse{
		writer: w,
	}
}

// Encode はデータをJSON形式にエンコードします
func (r *JSONResponse) Encode(statusCode int, data interface{}) error {
	r.writer.WriteHeader(statusCode)
	return json.NewEncoder(r.writer).Encode(data)
}

// CharsetMiddleware は文字コードを処理するミドルウェアです
func CharsetMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

// SanitizeString は文字列を正規化します
func SanitizeString(s string) string {
	if !utf8.ValidString(s) {
		v := make([]rune, 0, len(s))
		for i, r := range s {
			if r == utf8.RuneError {
				_, size := utf8.DecodeRuneInString(s[i:])
				if size == 1 {
					continue
				}
			}
			v = append(v, r)
		}
		s = string(v)
	}
	return s
} 