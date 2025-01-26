// ルーティングを定義するファイル
package routes

import (
	"employee-management/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

// マルチプレクサにエンドポイントとハンドラを登録する関数
func RegisterRoutes() *mux.Router {
	// 新しいルーターを作成
	router := mux.NewRouter()

	// エンドポイントとそのハンドラを登録
	router.HandleFunc("/employees", controllers.GetAllEmployees).Methods(http.MethodGet)        // 社員情報の全取得
	router.HandleFunc("/employees", controllers.AddEmployee).Methods(http.MethodPost)           // 新しい社員を登録
	router.HandleFunc("/employees", controllers.UpdateEmployee).Methods(http.MethodPut)         // 社員情報を更新
	router.HandleFunc("/employees/{id}", controllers.DeleteEmployee).Methods(http.MethodDelete) // 社員情報を削除

	// 他のエンドポイントをここに追加可能
	// 例: router.HandleFunc("/employees/{id}", controllers.UpdateEmployee).Methods(http.MethodPut)

	return router
}
