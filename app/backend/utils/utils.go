// 共通処理定義
package utils

import (
	"encoding/json"
	"net/http"
)

// エラー時のレスポンスをJSON形式で返却する関数
func ErrorResponse(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "appkication/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

// 正常時のレスポンスをJSON形式で返却する関数
func JSONRespose(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
