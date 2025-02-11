import { API_URL } from "./constants.js";
import { navigateToUpdatePage } from "./update.js";
import { deleteSelectedEmployee } from "./delete.js";

// 社員情報一覧ページの初期処理
export function setupEmployeePage() {
  fetchEmployeesForSelection();
  document.getElementById("update-button").addEventListener("click", navigateToUpdatePage);
  document.getElementById("delete-button").addEventListener("click", deleteSelectedEmployee);
}

// 社員情報の取得・表示処理
export async function fetchEmployeesForSelection() {
  try {
    const response = await fetch(`${API_URL}/employees`);
    if (!response.ok) throw new Error(`HTTPエラー: ${response.status}`);

    const employees = await response.json();

    if (!employees || !Array.isArray(employees)) {
      console.warn("無効なレスポンス形式: 配列を期待していましたが、受け取ったのは:", employees);
      return;
    }
    
    const tableBody = document.querySelector("#employee-table tbody");
    tableBody.innerHTML = "";

    if (employees.length === 0) {
      tableBody.innerHTML = `<tr><td colspan="6">社員情報が見つかりません。</td></tr>`;
      return;
    }

    employees.forEach((employee) => {
      const row = document.createElement("tr");
      row.innerHTML = `
        <td><input type="radio" name="selectedEmployee" value="${employee.id}"></td>
        <td>${employee.id}</td>
        <td>${employee.familyName}</td>
        <td>${employee.firstName}</td>
        <td>${employee.position}</td>
        <td>${employee.department}</td>
      `;
      tableBody.appendChild(row);
    });

  } catch (error) {
    console.error("社員情報の取得に失敗しました:", error);
    alert("社員情報の取得に失敗しました。");
  }
}
