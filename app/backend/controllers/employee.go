// ハンドラ関数を定義するファイル
package controllers

import (
	"employee-management/models"
	"employee-management/utils"
	"encoding/json"
	"net/http"
)

// 全社員情報を取得するハンドラ関数
//
// URL: GET /employees
// Response: 社員リストのJSON配列
func GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	// 社員情報取得関数実行
	employees, err := models.GetAllEmployees()
	if err != nil {
		utils.ErrorResponse(w, "社員情報の取得に失敗しました", http.StatusInternalServerError)
		return
	}

	// 取得結果をJSON形式に変換してレスポンス
	utils.JSONRespose(w, employees, http.StatusOK)
}

// 社員情報を登録するハンドラ関数
//
// URL: POST /employees
// Body: JSONオブジェクト
// Response: 成功メッセージ
func AddEmployee(w http.ResponseWriter, r *http.Request) {
	var emp models.Employee

	// Bodyをデコードし、Employee構造体に格納
	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		utils.ErrorResponse(w, "入力データが不正です", http.StatusBadRequest)
		return
	}

	// 社員情報登録関数実行
	err = models.AddEmployee(emp)
	if err != nil {
		utils.ErrorResponse(w, "社員情報の登録に失敗しました", http.StatusInternalServerError)
		return
	}

	// 実行結果をJSON形式に変換してレスポンス
	utils.JSONRespose(w, map[string]string{"message": "従業員の登録に成功しました"}, http.StatusOK)
}
