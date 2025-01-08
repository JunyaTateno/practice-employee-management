package main

import (
	"employee-management/config"
	"employee-management/routes"
	"log"
	"net/http"
)

func main() {
	// データベース接続の初期化
	config.ConnectDatabase()
	defer config.CloseDatabase()

	// ルーティングの設定
	router := routes.RegisterRoutes()

	// サーバー起動
	log.Println("サーバーを起動します :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
