// 定数定義
package constants

// CORSで許可するオリジン
var AllowedOrigins = map[string]bool{
	"http://localhost":                 true, // 開発環境 (Nginx)
	"https://your-production-site.com": true, // 本番環境 TODO: 本番環境のURLに変更する
}
