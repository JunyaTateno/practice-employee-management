// 定数定義
package constants

import "os"

var (
	DBUser     = getEnv("DB_USER", "youruser")         // DBユーザー名
	DBPassword = getEnv("DB_PASSWORD", "yourpassword") // DBパスワード
	DBHost     = getEnv("DB_HOST", "db")               // DBホスト名
	DBPort     = getEnv("DB_PORT", "3306")             // DBポート番号
	DBName     = getEnv("DB_NAME", "company")          // 使用するDB名
)

// 環境変数を取得するヘルパー関数
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// CORSで許可するオリジン
var AllowedOrigins = map[string]bool{
	"http://localhost":                 true, // 開発環境 (Nginx)
	"https://your-production-site.com": true, // 本番環境 TODO: 本番環境のURLに変更する
}
