package test

import (
	"bytes"
	"database/sql"
	"employee-management/config"
	"employee-management/models"
	"employee-management/routes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

// モックデータの定義
var mockEmployeeData = []models.Employee{
	{ID: 1, FamilyName: "Tanaka", FirstName: "Taro", Position: "Developer", Department: "IT"},
	{ID: 2, FamilyName: "Suzuki", FirstName: "Hanako", Position: "Manager", Department: "HR"},
}

// モックデータを生成する関数
func generateMockRows(data []models.Employee) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"id", "family_name", "first_name", "position", "department"})
	for _, emp := range data {
		rows.AddRow(emp.ID, emp.FamilyName, emp.FirstName, emp.Position, emp.Department)
	}
	return rows
}

// モックデータベースとルータをセットアップする関数
func setupMockDB() (*sql.DB, sqlmock.Sqlmock, http.Handler) {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	// アプリケーションのDB設定をモックDBに置き換える
	config.DB = db

	// ルータをセットアップ
	router := routes.RegisterRoutes()

	return db, mock, router
}

// HTTPリクエストを作成する共通関数
func makeRequest(t *testing.T, method, url string, body interface{}) (*httptest.ResponseRecorder, *http.Request) {
	var reqBody []byte
	if body != nil {
		var err error
		reqBody, err = json.Marshal(body)
		assert.NoError(t, err)
	}
	req := httptest.NewRequest(method, url, bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	return w, req
}

// テスト: 正常系- GetAllEmployees エンドポイント
func TestE2EGetAllEmployees(t *testing.T) {
	db, mock, router := setupMockDB()
	defer db.Close()

	// モックデータを生成
	selectQuery := "SELECT (.+) FROM employees"
	mockRows := generateMockRows(mockEmployeeData)
	mock.ExpectQuery(selectQuery).WillReturnRows(mockRows)

	// HTTPリクエスト
	w, req := makeRequest(t, http.MethodGet, "/employees", nil)
	router.ServeHTTP(w, req)

	// レスポンスの検証
	resp := w.Result()
	defer resp.Body.Close()

	// ステータスコードの検証
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// レスポンスボディの検証
	var employees []models.Employee
	err := json.NewDecoder(resp.Body).Decode(&employees)
	assert.NoError(t, err)

	// モックデータとレスポンスボディを比較
	assert.Equal(t, len(mockEmployeeData), len(employees), "従業員の数が一致しません")
	for i, expected := range mockEmployeeData {
		assert.Equal(t, expected, employees[i], "従業員データが一致しません")
	}

	// モッククエリの期待値が満たされているか確認
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

// テスト: 正常系- AddEmployee エンドポイント
func TestE2EAddEmployee(t *testing.T) {
	db, mock, router := setupMockDB()
	defer db.Close()

	// INSERT用の各項目
	familyName := "Kato"
	firstName := "Jiro"
	position := "Designer"
	department := "Design"

	// モッククエリの設定
	insertQuery := "INSERT INTO employees (.+)"
	mock.ExpectExec(insertQuery).
		WithArgs(familyName, firstName, position, department). // 変数を利用
		WillReturnResult(sqlmock.NewResult(3, 1))              // ID: 3, RowsAffected: 1

	// HTTPリクエストの準備
	// リクエストボディを作成
	newEmployee := models.Employee{
		FamilyName: familyName,
		FirstName:  firstName,
		Position:   position,
		Department: department,
	}

	// HTTPリクエスト
	w, req := makeRequest(t, http.MethodPost, "/employees", newEmployee)
	router.ServeHTTP(w, req)

	// レスポンスの検証
	resp := w.Result()
	defer resp.Body.Close()

	// ステータスコードの検証
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// レスポンスボディの検証
	var resBody map[string]string
	err := json.NewDecoder(resp.Body).Decode(&resBody)
	assert.NoError(t, err)

	// モッククエリの期待値が満たされているか確認
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

// テスト: 正常系 - UpdateEmployee エンドポイント
func TestE2EUpdateEmployee(t *testing.T) {
	db, mock, router := setupMockDB()
	defer db.Close()

	// 各項目を変数に定義
	id := 1
	updatedEmployee := models.Employee{
		ID:         id,
		FamilyName: "Yamada",
		FirstName:  "Taro",
		Position:   "Team Leader",
		Department: "Development",
	}

	// モッククエリ: 更新処理
	updateQuery := "UPDATE employees SET (.+) WHERE id = ?"
	mock.ExpectExec(updateQuery).
		WithArgs(updatedEmployee.FamilyName, updatedEmployee.FirstName, updatedEmployee.Position, updatedEmployee.Department, id).
		WillReturnResult(sqlmock.NewResult(0, 1)) // RowsAffected: 1

	// モッククエリ: 更新後のデータを取得
	mockRows := generateMockRows([]models.Employee{updatedEmployee})
	selectQuery := "SELECT (.+) FROM employees WHERE id = ?"
	mock.ExpectQuery(selectQuery).
		WithArgs(id).
		WillReturnRows(mockRows)

	// HTTPリクエスト (`PUT /employees/{id}`)
	url := fmt.Sprintf("/employees/%d", id)
	w, req := makeRequest(t, http.MethodPut, url, updatedEmployee)
	router.ServeHTTP(w, req)

	// レスポンスの検証
	resp := w.Result()
	defer resp.Body.Close()

	// ステータスコードの検証
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// レスポンスボディの検証
	var resBody models.Employee
	err := json.NewDecoder(resp.Body).Decode(&resBody)
	assert.NoError(t, err)

	// レスポンスが更新後のデータと一致することを確認
	assert.Equal(t, id, resBody.ID)
	assert.Equal(t, updatedEmployee.FamilyName, resBody.FamilyName)
	assert.Equal(t, updatedEmployee.FirstName, resBody.FirstName)
	assert.Equal(t, updatedEmployee.Position, resBody.Position)
	assert.Equal(t, updatedEmployee.Department, resBody.Department)

	// モッククエリの期待値が満たされているか確認
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

// テスト: 正常系- DeleteEmployee エンドポイント
func TestE2EDeleteEmployee(t *testing.T) {
	db, mock, router := setupMockDB()
	defer db.Close()

	// 各項目を変数に定義
	id := 1

	// モッククエリの設定
	deleteQuery := "DELETE FROM employees WHERE id = ?"
	mock.ExpectExec(deleteQuery).
		WithArgs(id).
		WillReturnResult(sqlmock.NewResult(0, 1)) // RowsAffected: 1

	// HTTPリクエストの準備
	w, req := makeRequest(t, http.MethodDelete, "/employees/1", nil)
	router.ServeHTTP(w, req)

	// レスポンスの検証
	resp := w.Result()
	defer resp.Body.Close()

	// ステータスコードの検証
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// レスポンスボディの検証
	var resBody map[string]string
	err := json.NewDecoder(resp.Body).Decode(&resBody)
	assert.NoError(t, err)

	// モッククエリの期待値が満たされているか確認
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
