// 社員テーブルへの操作
package models

import (
	"employee-management/config"
)

// Employee構造体の定義
type Employee struct {
	ID         int    `json:"id"`
	FamilyName string `json:"familyName"`
	FirstName  string `json:"firstName"`
	Position   string `json:"position"`
	Department string `json:"department"`
}

// DBから全社員情報を取得する関数
func GetAllEmployees() ([]Employee, error) {
	// SELECTクエリ実行
	rows, err := config.DB.Query("SELECT * FROM employees")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Employee構造体スライスに格納
	var employees []Employee
	for rows.Next() {
		var emp Employee
		err := rows.Scan(&emp.ID, &emp.FamilyName, &emp.FirstName, &emp.Position, &emp.Department)
		if err != nil {
			return nil, err
		}
		employees = append(employees, emp)
	}

	return employees, nil
}

// DBに社員情報をインサートする関数
func AddEmployee(emp Employee) error {
	_, err := config.DB.Exec("INSERT INTO employees (id) VALUES (?)",
		emp.ID)
	return err
}
