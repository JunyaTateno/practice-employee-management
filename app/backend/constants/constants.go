// 定数定義
package constants

// ファイルパス
var (
	APP_ROOT   = "/app"
	QUERY_ROOT = "queries"
)

// SQLファイル名
var (
	SELECT_QUERY       = "select_employee.sql"
	SELECT_BY_ID_QUERY = "select_employee_by_id.sql"
	INSERT_QUERY       = "insert_employee.sql"
	UPDATE_QUERY       = "update_employee.sql"
	DELETE_QUERY       = "delete_employee.sql"
)

// CORSで許可するオリジン
var ALLOWED_ORIGINS = map[string]bool{
	"http://localhost":    true, // 開発環境 (Nginx)
	"http://35.75.63.143": true, // 本番環境
}

// 社員名の最大文字数
const NAME_MAX_LENGTH = 30

// 役職のリスト
var VALID_POSITIONS = []string{"一般社員", "主任", "課長", "部長"}

// 部署のリスト
var VALID_DEPARTMENTS = []string{"技術部", "経理部", "人事部", "営業部"}
