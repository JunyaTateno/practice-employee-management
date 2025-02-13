// 社員情報のモデルを定義するファイル
package models

import (
	"database/sql"
	"employee-management/config"
	"employee-management/constants"
	"employee-management/queries"
	"errors"
	"slices"
)

// Employee構造体の定義
type Employee struct {
	ID         int    `json:"id"`
	FamilyName string `json:"familyName"`
	FirstName  string `json:"firstName"`
	Position   string `json:"position"`
	Department string `json:"department"`
}

// バリデーション関数
func ValidateEmployee(emp Employee) error {
	// 必須項目チェック
	if emp.FamilyName == "" || emp.FirstName == "" {
		return errors.New("姓・名は必須です")
	}

	// 文字数制限 (最大30文字)
	if len(emp.FamilyName) > constants.NAME_MAX_LENGTH || len(emp.FirstName) > constants.NAME_MAX_LENGTH {
		return errors.New("姓・名は30文字以内で入力してください")
	}

	// 役職がリストに含まれているか
	if !slices.Contains(constants.VALID_POSITIONS, emp.Position) {
		return errors.New("無効な役職が指定されました")
	}

	// 部署がリストに含まれているか
	if !slices.Contains(constants.VALID_DEPARTMENTS, emp.Department) {
		return errors.New("無効な部署が指定されました")
	}

	return nil
}

// DBから全社員情報を取得する関数
func GetAllEmployees() ([]Employee, error) {
	// キャッシュからSQLを取得
	selectSQL, err := queries.GetSQL(constants.SELECT_QUERY)
	if err != nil {
		return nil, err
	}

	// SELECTクエリ実行
	rows, err := config.DB.Query(selectSQL)
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
	// キャッシュからSQLを取得
	insertSQL, err := queries.GetSQL(constants.INSERT_QUERY)
	if err != nil {
		return err
	}

	_, err = config.DB.Exec(insertSQL, emp.FamilyName, emp.FirstName, emp.Position, emp.Department)
	return err
}

// DBの社員情報を更新し、変更後のデータを返す関数
func UpdateEmployee(emp Employee) (Employee, error) {
	// キャッシュからSQLを取得
	updateSQL, err := queries.GetSQL(constants.UPDATE_QUERY)
	if err != nil {
		return Employee{}, err
	}
	selectSQL, err := queries.GetSQL(constants.SELECT_BY_ID_QUERY)
	if err != nil {
		return Employee{}, err
	}

	// トランザクション開始
	tx, err := config.DB.Begin()
	if err != nil {
		return Employee{}, err
	}

	// UPDATE 文を実行
	result, err := tx.Exec(updateSQL, emp.FamilyName, emp.FirstName, emp.Position, emp.Department, emp.ID)

	if err != nil {
		tx.Rollback()
		return Employee{}, err
	}

	// 更新対象が存在しない場合
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return Employee{}, err
	}
	if rowsAffected == 0 {
		tx.Rollback()
		return Employee{}, sql.ErrNoRows
	}

	// 更新後のデータを取得
	var updatedEmp Employee
	err = tx.QueryRow(selectSQL, emp.ID).Scan(&updatedEmp.ID, &updatedEmp.FamilyName, &updatedEmp.FirstName, &updatedEmp.Position, &updatedEmp.Department)
	if err != nil {
		tx.Rollback()
		return Employee{}, err
	}

	// トランザクションをコミット
	err = tx.Commit()
	if err != nil {
		return Employee{}, err
	}

	return updatedEmp, nil
}

// 社員情報を削除する関数
func DeleteEmployee(id int) error {
	// キャッシュからSQLを取得
	deleteSQL, err := queries.GetSQL(constants.DELETE_QUERY)
	if err != nil {
		return err
	}

	// トランザクション開始
	tx, err := config.DB.Begin()
	if err != nil {
		return err
	}

	// DELETE 文を実行
	result, err := tx.Exec(deleteSQL, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	// 削除対象が存在しない場合
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}
	if rowsAffected == 0 {
		tx.Rollback()
		return sql.ErrNoRows
	}

	// トランザクションをコミット
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
