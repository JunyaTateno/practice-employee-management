// ハンドラ関数を定義するファイル
package controllers

import (
	"employee-management/models"
	"employee-management/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

	// BodyをJSONとしてデコードし、Employee構造体に格納
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

// 社員情報を更新するハンドラ関数
//
// URL: PUT /employees/{id}
// Body: JSONオブジェクト
// Response: 成功メッセージ
func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	var emp models.Employee

	// Body を JSON としてデコードし、Employee 構造体に格納
	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		utils.ErrorResponse(w, "入力データが不正です", http.StatusBadRequest)
		return
	}

	// 社員情報更新関数実行
	updatedEmployee, err := models.UpdateEmployee(emp)
	if err != nil {
		utils.ErrorResponse(w, "社員情報の更新に失敗しました", http.StatusInternalServerError)
		return
	}

	// 更新後の社員情報をレスポンスとして返す
	utils.JSONRespose(w, updatedEmployee, http.StatusOK)
}

// 社員情報を削除するハンドラ関数
//
// URL: DELETE /employees/{id}
// Response: 成功メッセージ
func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	// URL パスから ID を取得
	idStr := mux.Vars(r)["id"]

	// ID を数値に変換
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.ErrorResponse(w, "IDが無効です", http.StatusBadRequest)
		return
	}

	// 社員情報削除関数を実行
	err = models.DeleteEmployee(id)
	if err != nil {
		utils.ErrorResponse(w, "社員情報の削除に失敗しました", http.StatusInternalServerError)
		return
	}

	// 成功レスポンス
	utils.JSONRespose(w, map[string]string{"message": "従業員の削除に成功しました"}, http.StatusOK)
}
