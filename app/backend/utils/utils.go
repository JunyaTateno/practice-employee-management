// 共通処理関数を定義するファイル
package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

// 環境変数を取得する関数
func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// エラー時のレスポンスをJSON形式で返却する関数
func ErrorResponse(w http.ResponseWriter, message string, err error, status int) {
	log.Printf("Error: %s: %v", message, err)
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
