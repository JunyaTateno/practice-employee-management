import { API_URL } from "./constants.js";
import { fetchEmployeesForSelection } from "./employee.js";
import { toggleActionButtons } from "./utils.js";

// 選択された社員を削除する
export async function deleteSelectedEmployee() {
  const selectedEmployee = document.querySelector("input[name='selectedEmployee']:checked");
  
  if (!selectedEmployee) {
    alert("削除する社員を選択してください。");
    return;
  }

  const employeeId = parseInt(selectedEmployee.value, 10);
  if (isNaN(employeeId)) {
    alert("有効な社員IDを取得できませんでした。");
    return;
  }

  if (!confirm("本当にこの社員を削除しますか？")) {
    return;
  }

  try {
    const response = await fetch(`${API_URL}/employees/${employeeId}`, {
      method: "DELETE",
    });

    if (response.ok) {
      alert("社員が削除されました！");
      fetchEmployeesForSelection(); // 削除後にリストを更新
      toggleActionButtons(); // ボタンの状態を更新
    } else {
      throw new Error("社員の削除に失敗しました。");
    }
  } catch (error) {
    console.error("社員の削除中にエラーが発生しました:", error);
    alert("社員の削除に失敗しました。もう一度試してください。");
  }
}