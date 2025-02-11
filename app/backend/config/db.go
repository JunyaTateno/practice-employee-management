// DB接続を管理するファイル
package config

import (
	"database/sql"
	"employee-management/utils"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// DB接続プール用変数
var DB *sql.DB

// 環境変数から取得するDB設定
var (
	DBUser     = utils.GetEnv("DB_USER", "youruser")         // DBユーザー名
	DBPassword = utils.GetEnv("DB_PASSWORD", "yourpassword") // DBパスワード
	DBHost     = utils.GetEnv("DB_HOST", "db")               // DBホスト名
	DBPort     = utils.GetEnv("DB_PORT", "3306")             // DBポート番号
	DBName     = utils.GetEnv("DB_NAME", "company")          // 使用するDB名
)

// データベース接続を初期化する関数
func ConnectDatabase() {
	var err error

	// MySQL接続文字列
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?timeout=10s",
		DBUser,
		DBPassword,
		DBHost,
		DBPort,
		DBName,
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
		log.Println("DB接続のクローズに失敗しました:", err)
	} else {
		log.Println("DB接続のクローズに成功しました")
	}
}
