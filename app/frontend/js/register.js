import { API_URL } from "./constants.js";
import { populateDropdowns, validateNameLength, validatePosition, validateDepartment } from "./utils.js";

// 登録フォームの処理
export function setupRegisterForm() {
  const form = document.getElementById("employee-form");
  populateDropdowns();

  form.addEventListener("submit", async (event) => {
    event.preventDefault();

    const familyName = document.getElementById("familyName").value;
    const firstName = document.getElementById("firstName").value;
    const position = document.getElementById("position").value;
    const department = document.getElementById("department").value;

    // バリデーションチェック (失敗したら処理を中断)
    if (!validateNameLength(familyName, "姓") || 
        !validateNameLength(firstName, "名") || 
        !validatePosition(position) || 
        !validateDepartment(department)) {
      return;
    }

    const newEmployee = { familyName, firstName, position, department };

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
        headers: { "Content-Type": "application/json" },
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
