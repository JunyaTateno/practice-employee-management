const API_URL = "http://localhost:8080"; // バックエンドのURL

// DOMContentLoaded イベントで初期化
document.addEventListener("DOMContentLoaded", () => {
  const currentPage = document.body.id;

  switch (currentPage) {
    case "view-page":
      fetchEmployees();
      break;
    case "register-page":
      setupRegisterForm();
      break;
    case "update-page":
      setupUpdateForm();
      break;
    case "delete-page":
      setupDeleteForm();
      break;
    default:
      console.log("特定のページ機能はありません。");
  }
});

// 1. 参照ページの処理
async function fetchEmployees() {
  try {
    const response = await fetch(`${API_URL}/employees`);
    console.log("HTTPレスポンスステータス:", response.status);

    if (!response.ok) {
      throw new Error(`HTTPエラー！ステータス: ${response.status}`);
    }

    let employees = await response.json(); // `let` に変更
    console.log("APIレスポンス JSON:", employees);

    if (!employees || !Array.isArray(employees)) {
      console.warn("無効なレスポンス形式: 配列を期待していましたが、受け取ったのは:", employees);
      employees = []; // `null` の場合は `[]` に置き換え
    }

    const tableBody = document.querySelector("#employee-table tbody");
    tableBody.innerHTML = "";

    if (employees.length === 0) {
      tableBody.innerHTML = `
        <tr>
          <td colspan="5">社員情報が見つかりません。</td>
        </tr>
      `;
      return;
    }

    employees.forEach((employee) => {
      const row = document.createElement("tr");
      row.innerHTML = `
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
    alert("社員情報の取得に失敗しました。詳細はコンソールを確認してください。");
  }
}

// 2. 登録ページの処理
function setupRegisterForm() {
  const form = document.getElementById("employee-form");
  form.addEventListener("submit", async (event) => {
    event.preventDefault();

    const newEmployee = {
      familyName: document.getElementById("familyName").value,
      firstName: document.getElementById("firstName").value,
      position: document.getElementById("position").value,
      department: document.getElementById("department").value,
    };

    // 確認ダイアログを表示
    const confirmationMessage = `以下の情報で登録しますか？\n\n` +
      `姓: ${newEmployee.familyName}\n` +
      `名: ${newEmployee.firstName}\n` +
      `役職: ${newEmployee.position}\n` +
      `部署: ${newEmployee.department}`;
    
    if (!confirm(confirmationMessage)) {
      return;
    }

    console.log("🚀 新しい社員情報を送信中:", newEmployee);

    try {
      const response = await fetch(`${API_URL}/employees`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(newEmployee),
      });

      console.log("📡 HTTPレスポンスステータス:", response.status);

      if (!response.ok) {
        const errorText = await response.text();
        throw new Error(`HTTPエラー！ステータス: ${response.status}, レスポンス: ${errorText}`);
      }

      alert("社員の登録が完了しました！");
      form.reset();
    } catch (error) {
      console.error("❌ 社員の登録に失敗しました:", error);
      alert("社員の登録に失敗しました。詳細はコンソールを確認してください。");
    }
  });
}

// 3. 更新ページの処理
function setupUpdateForm() {
  fetchEmployeesForSelection("update");
  document.getElementById("update-button").addEventListener("click", updateSelectedEmployee);
}

// 4. 削除ページの処理
function setupDeleteForm() {
  fetchEmployeesForSelection("delete");
  document.getElementById("delete-button").addEventListener("click", deleteSelectedEmployee);
}  

// 全社員情報を取得し、ラジオボタン付きのリストを表示
async function fetchEmployeesForSelection(action) {
  try {
    const response = await fetch(`${API_URL}/employees`);
    console.log("HTTPレスポンスステータス:", response.status);

    if (!response.ok) {
      throw new Error(`HTTPエラー！ステータス: ${response.status}`);
    }

    const employees = await response.json();
    console.log("APIレスポンス JSON:", employees);

    if (!employees || !Array.isArray(employees)) {
      console.warn("無効なレスポンス形式: 配列を期待していましたが、受け取ったのは:", employees);
      return;
    }

    const tableBody = document.querySelector("#employee-table tbody");
    if (!tableBody) {
      console.error("エラー: #employee-table tbody が見つかりません。");
      return;
    }
    
    tableBody.innerHTML = "";

    if (employees.length === 0) {
      tableBody.innerHTML = `
        <tr>
          <td colspan="6">社員情報が見つかりません。</td>
        </tr>
      `;
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
    alert("社員情報の取得に失敗しました。詳細はコンソールを確認してください。");
  }
}

// ラジオボタンで選択された社員の情報をフォームに反映
document.addEventListener("change", (event) => {
  if (event.target.name === "selectedEmployee") {
    const selectedEmployee = event.target.closest("tr");
    document.getElementById("employeeId").value = event.target.value;
    document.getElementById("familyName").value = selectedEmployee.children[2].textContent;
    document.getElementById("firstName").value = selectedEmployee.children[3].textContent;
    document.getElementById("position").value = selectedEmployee.children[4].textContent;
    document.getElementById("department").value = selectedEmployee.children[5].textContent;
  }
});

// 選択された社員の情報を更新
async function updateSelectedEmployee(event) {
  event.preventDefault();

  // `parseInt()` で整数変換し、NaN の場合は削除する
  const employeeId = parseInt(document.getElementById("employeeId").value, 10);
  if (isNaN(employeeId)) {
    alert("更新する社員を選択してください。");
    return;
  }

  const updatedEmployee = {
    id: employeeId,
    familyName: document.getElementById("familyName").value || null,
    firstName: document.getElementById("firstName").value || null,
    position: document.getElementById("position").value || null,
    department: document.getElementById("department").value || null,
  };

  // 確認ダイアログを表示
  const confirmationMessage = `以下の情報で更新しますか？\n\n` +
    `ID: ${updatedEmployee.id}\n` +
    `姓: ${updatedEmployee.familyName}\n` +
    `名: ${updatedEmployee.firstName}\n` +
    `役職: ${updatedEmployee.position}\n` +
    `部署: ${updatedEmployee.department}`;
  
  if (!confirm(confirmationMessage)) {
    return;
  }

  try {
    const response = await fetch(`${API_URL}/employees`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(updatedEmployee),
    });

    if (response.ok) {
      alert("社員情報が更新されました！");
      fetchEmployeesForSelection("update"); // 更新後にリストを更新
    } else {
      throw new Error("社員情報の更新に失敗しました。");
    }
  } catch (error) {
    console.error("社員情報の更新中にエラーが発生しました:", error);
    alert("社員情報の更新に失敗しました。もう一度試してください。");
  }
}

// 選択された社員を削除
async function deleteSelectedEmployee() {
  const selectedEmployee = document.querySelector("input[name='selectedEmployee']:checked");
  if (!selectedEmployee) {
    alert("削除する社員を選択してください。");
    return;
  }

  const employeeId = selectedEmployee.value;
  if (!confirm(`社員ID: ${employeeId} を削除しますか？`)) {
    return;
  }

  try {
    const response = await fetch(`${API_URL}/employees/${employeeId}`, {
      method: "DELETE",
    });

    if (response.ok) {
      alert("社員情報が削除されました！");
      fetchEmployeesForSelection("delete"); // 削除後にリストを更新
    } else {
      throw new Error("社員情報の削除に失敗しました。");
    }
  } catch (error) {
    console.error("社員情報の削除中にエラーが発生しました:", error);
    alert("社員情報の削除に失敗しました。もう一度試してください。");
  }
}
