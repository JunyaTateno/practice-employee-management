package queries

import (
	"employee-management/constants"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

// SQLクエリのキャッシュ用マップ
var (
	sqlQueries = make(map[string]string)
	mu         sync.RWMutex
)

// アプリ起動時にSQLを一度だけ読み込む
func init() {
	sqlFiles := []string{
		constants.SelectQuery,
		constants.InsertQuery,
		constants.UpdateQuery,
		constants.DeleteQuery}

	for _, fileName := range sqlFiles {
		filePath := filepath.Join("queries", fileName)
		sqlBytes, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Printf("SQLファイルの読み込みに失敗しました: %s, エラー: %v\n", fileName, err)
			continue
		}

		// マップに保存
		sqlQueries[fileName] = string(sqlBytes)
	}
}

// キャッシュされたSQLを取得する
func GetSQL(fileName string) (string, error) {
	mu.RLock()
	defer mu.RUnlock()

	query, exists := sqlQueries[fileName]
	if !exists {
		return "", fmt.Errorf("指定されたSQLファイルが見つかりません: %s", fileName)
	}
	return query, nil
}
