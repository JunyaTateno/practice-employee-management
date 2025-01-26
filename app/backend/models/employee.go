// 社員情報のモデルを定義するファイル
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
	rows, err := config.DB.Query("SELECT id, family_name, first_name, position, department FROM employees")
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
	_, err := config.DB.Exec("INSERT INTO employees (family_name, first_name, position, department) VALUES (?, ?, ?, ?)",
		emp.FamilyName, emp.FirstName, emp.Position, emp.Department)
	return err
}

// DBの社員情報を更新し、変更後のデータを返す関数
func UpdateEmployee(emp Employee) (Employee, error) {
	// UPDATE 文を実行
	_, err := config.DB.Exec(`
        UPDATE employees
        SET family_name = ?, first_name = ?, position = ?, department = ?
        WHERE id = ?
    `, emp.FamilyName, emp.FirstName, emp.Position, emp.Department, emp.ID)
	if err != nil {
		return Employee{}, err
	}

	// 更新後のデータを取得
	var updatedEmp Employee
	err = config.DB.QueryRow(`
        SELECT id, family_name, first_name, position, department
        FROM employees
        WHERE id = ?
    `, emp.ID).Scan(&updatedEmp.ID, &updatedEmp.FamilyName, &updatedEmp.FirstName, &updatedEmp.Position, &updatedEmp.Department)
	if err != nil {
		return Employee{}, err
	}

	return updatedEmp, nil
}

// 社員情報を削除する関数
func DeleteEmployee(id int) error {
	query := `DELETE FROM employees WHERE id = ?`

	_, err := config.DB.Exec(query, id)
	return err
}
