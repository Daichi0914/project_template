package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewJSONResponse(t *testing.T) {
	// テスト用のレスポンスレコーダーを作成
	w := httptest.NewRecorder()
	
	// JSONResponseの作成
	resp := NewJSONResponse(w)
	
	// 検証
	assert.NotNil(t, resp)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
}

func TestJSONResponse_Encode(t *testing.T) {
	// テスト用のレスポンスレコーダーを作成
	w := httptest.NewRecorder()
	
	// JSONResponseの作成
	resp := NewJSONResponse(w)
	
	// テストデータ
	testData := map[string]string{
		"message": "テストメッセージ",
	}
	
	// エンコード実行
	err := resp.Encode(http.StatusOK, testData)
	
	// 検証
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "テストメッセージ")
}

func TestSanitizeString_ValidString(t *testing.T) {
	// 正常な文字列のテスト
	input := "こんにちは世界"
	result := SanitizeString(input)
	
	// 検証 - 正常な文字列は変更されないはず
	assert.Equal(t, input, result)
}

func TestSanitizeString_InvalidString(t *testing.T) {
	// 不正なUTF-8文字を含む文字列
	// 0xE0 0x80 はUTF-8として不正なバイト列
	invalidBytes := []byte{0xE0, 0x80, 0xAF}
	input := string(invalidBytes) + "こんにちは世界"
	
	// サニタイズ実行
	result := SanitizeString(input)
	
	// 検証 - 不正な文字が除去されているはず
	assert.NotEqual(t, input, result)
	assert.Contains(t, result, "こんにちは世界")
}

func TestCharsetMiddleware(t *testing.T) {
	// テスト用のハンドラー
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	
	// ミドルウェアの作成
	middleware := CharsetMiddleware(testHandler)
	
	// テスト用のリクエストとレスポンスレコーダー
	req := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()
	
	// ミドルウェア実行
	middleware.ServeHTTP(w, req)
	
	// 検証
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "OK", w.Body.String())
} 