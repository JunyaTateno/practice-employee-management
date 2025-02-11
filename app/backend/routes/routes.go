// ルーティングを定義するファイル
package routes

import (
	"employee-management/constants"
	"employee-management/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

// CORSミドルウェアを定義
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		// 許可されたオリジンの場合、CORSヘッダーを追加
		if constants.AllowedOrigins[origin] {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		}

		// `OPTIONS` メソッドのリクエストには即時レスポンスを返す
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// 通常のリクエストも CORS ヘッダーを持つようにする
		next.ServeHTTP(w, r)
	})
}

// マルチプレクサにエンドポイントとハンドラを登録する関数
func RegisterRoutes() http.Handler {
	// 新しいルーターを作成
	router := mux.NewRouter()

	// エンドポイントとそのハンドラを登録
	router.HandleFunc("/employees", controllers.GetAllEmployees).Methods(http.MethodGet)        // 社員情報の全取得
	router.HandleFunc("/employees", controllers.AddEmployee).Methods(http.MethodPost)           // 新しい社員を登録
	router.HandleFunc("/employees/{id}", controllers.UpdateEmployee).Methods(http.MethodPut)    // 社員情報を更新
	router.HandleFunc("/employees/{id}", controllers.DeleteEmployee).Methods(http.MethodDelete) // 社員情報を削除

	// 他のエンドポイントをここに追加可能
	// 例: router.HandleFunc("/employees/{id}", controllers.UpdateEmployee).Methods(http.MethodPut)

	// CORS ミドルウェアを適用
	return enableCORS(router)
}
