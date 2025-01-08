// DB接続を管理するファイル
package config

import (
	"database/sql"
	"employee-management/constants"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// DB接続プール用変数
var DB *sql.DB

// データベース接続を初期化する関数
func ConnectDatabase() {
	var err error

	// MySQL接続文字列
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?timeout=10s",
		constants.DBUser,
		constants.DBPassword,
		constants.DBHost,
		constants.DBPort,
		constants.DBName,
	)

	// DB接続オープン
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("DB接続の初期化に失敗しました:", err)
	}

	// DB接続確認
	err = DB.Ping()
	if err != nil {
		log.Fatal("DB接続確認できません:", err)
	}

	log.Println("DB接続に成功しました")
}

// データベース接続を終了する関数
func CloseDatabase() {
	err := DB.Close()
	if err != nil {
		log.Println("DB接続のクローズに失敗しました")
	} else {
		log.Println("DB接続のクローズに成功しました")
	}
}
