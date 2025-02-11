import { API_URL } from "./constants.js";
import { populateDropdowns } from "./utils.js";

// 変更ページに遷移する前に、選択された社員の情報を保存
export function navigateToUpdatePage() {
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
export function loadUpdateForm() {
  const employeeData = JSON.parse(localStorage.getItem("selectedEmployee"));
  if (!employeeData) {
    alert("無効なアクセスです。社員が選択されていません。");
    window.location.href = "index.html";
    return;
  }

  document.getElementById("employeeId").value = employeeData.id;
  document.getElementById("familyName").value = employeeData.familyName;
  document.getElementById("firstName").value = employeeData.firstName;
  populateDropdowns(employeeData.position, employeeData.department);
}

// 更新処理のセットアップ
export function setupUpdateForm() {
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
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(updatedEmployee),
      });

      if (response.ok) {
        alert("社員情報が更新されました！");
        localStorage.removeItem("selectedEmployee");
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
