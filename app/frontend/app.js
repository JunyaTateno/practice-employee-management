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
  const form = document.getElementById("update-form");
  form.addEventListener("submit", async (event) => {
    event.preventDefault();

    // `parseInt()` で整数変換し、NaN の場合は削除する
    const employeeId = parseInt(document.getElementById("employeeId").value, 10);
    if (isNaN(employeeId)) {
      alert("有効な社員IDを入力してください。");
      return;
    }

    const updatedEmployee = {
      id: employeeId, // `undefined` にしない
      familyName: document.getElementById("familyName").value || null,
      firstName: document.getElementById("firstName").value || null,
      position: document.getElementById("position").value || null,
      department: document.getElementById("department").value || null,
    };

    // 空の値（`null`）を削除
    Object.keys(updatedEmployee).forEach((key) => {
      if (updatedEmployee[key] === null) delete updatedEmployee[key];
    });

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
        form.reset();
      } else {
        throw new Error("社員情報の更新に失敗しました。");
      }
    } catch (error) {
      console.error("社員情報の更新中にエラーが発生しました:", error);
      alert("社員情報の更新に失敗しました。もう一度試してください。");
    }
  });
}

// 4. 削除ページの処理
function setupDeleteForm() {
  const form = document.getElementById("delete-form");
  form.addEventListener("submit", async (event) => {
    event.preventDefault();

    const employeeId = document.getElementById("employeeId").value;

    try {
      const response = await fetch(`${API_URL}/employees/${employeeId}`, {
        method: "DELETE",
      });

      if (response.ok) {
        alert("社員情報が削除されました！");
        form.reset();
      } else {
        throw new Error("社員情報の削除に失敗しました。");
      }
    } catch (error) {
      console.error("社員情報の削除中にエラーが発生しました:", error);
      alert("社員情報の削除に失敗しました。もう一度試してください。");
    }
  });
}
