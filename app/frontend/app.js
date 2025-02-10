const API_URL = "http://localhost:8080"; // バックエンドのURL

document.addEventListener("DOMContentLoaded", () => {
  const pageId = document.body.id;

  if (pageId === "employee-management") {
    setupEmployeePage();
  } else if (pageId === "register-page") {
    setupRegisterForm();
  } else if (pageId === "update-page") {
    loadUpdateForm();
    setupUpdateForm();
  }  
});  

// 参照・変更・削除をまとめたページの処理
function setupEmployeePage() {
  fetchEmployeesForSelection();
  document.getElementById("update-button").addEventListener("click", navigateToUpdatePage);
  document.getElementById("delete-button").addEventListener("click", deleteSelectedEmployee);
}  

// 全社員情報を取得し、ラジオボタン付きのリストを表示
async function fetchEmployeesForSelection() {
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

// 社員登録処理のセットアップ
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

    if (!confirm("以下の内容で登録しますか？\n\n" +
      `姓: ${newEmployee.familyName}\n` +
      `名: ${newEmployee.firstName}\n` +
      `役職: ${newEmployee.position}\n` +
      `部署: ${newEmployee.department}`)) {
      return;  
    }  

    try {
      const response = await fetch(`${API_URL}/employees`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },  
        body: JSON.stringify(newEmployee),
      });  

      if (response.ok) {
        alert("社員が登録されました！");
        form.reset();
      } else {
        throw new Error("社員の登録に失敗しました。");
      }  
    } catch (error) {
      console.error("社員の登録中にエラーが発生しました:", error);
      alert("社員の登録に失敗しました。もう一度試してください。");
    }  
  });  
}  

// 変更ページに遷移する前に、選択された社員の情報を保存
function navigateToUpdatePage() {
  const selectedEmployee = document.querySelector("input[name='selectedEmployee']:checked");
  if (!selectedEmployee) {
    alert("変更する社員を選択してください。");
    return;
  }

  const employeeRow = selectedEmployee.closest("tr").children;
  const employeeData = {
    id: selectedEmployee.value,
    familyName: employeeRow[2].textContent,
    firstName: employeeRow[3].textContent,
    position: employeeRow[4].textContent,
    department: employeeRow[5].textContent,
  };

  localStorage.setItem("selectedEmployee", JSON.stringify(employeeData));
  window.location.href = "update.html";
}

// 変更ページで社員情報をロード
function loadUpdateForm() {
  const employeeData = JSON.parse(localStorage.getItem("selectedEmployee"));
  if (!employeeData) {
    alert("無効なアクセスです。社員が選択されていません。");
    window.location.href = "index.html";
    return;
  }

  document.getElementById("employeeId").value = employeeData.id;
  document.getElementById("familyName").value = employeeData.familyName;
  document.getElementById("firstName").value = employeeData.firstName;
  document.getElementById("position").value = employeeData.position;
  document.getElementById("department").value = employeeData.department;
}

// 更新処理のセットアップ
function setupUpdateForm() {
  document.getElementById("update-form").addEventListener("submit", async (event) => {
    event.preventDefault();

    const updatedEmployee = {
      id: parseInt(document.getElementById("employeeId").value, 10),
      familyName: document.getElementById("familyName").value,
      firstName: document.getElementById("firstName").value,
      position: document.getElementById("position").value,
      department: document.getElementById("department").value,
    };

    if (!confirm("社員情報を更新しますか？")) {
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
        localStorage.removeItem("selectedEmployee"); // 更新後にキャッシュを削除
        window.location.href = "index.html";
      } else {
        throw new Error("社員情報の更新に失敗しました。");
      }
    } catch (error) {
      console.error("社員情報の更新中にエラーが発生しました:", error);
      alert("社員情報の更新に失敗しました。もう一度試してください。");
    }
  });
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
      fetchEmployeesForSelection();
    } else {
      throw new Error("社員情報の削除に失敗しました。");
    }
  } catch (error) {
    console.error("社員情報の削除中にエラーが発生しました:", error);
    alert("社員情報の削除に失敗しました。もう一度試してください。");
  }
}