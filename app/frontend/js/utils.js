import { positions, departments, NAME_MAX_LENGTH } from "./constants.js";

// ボタンの有効化・無効化を切り替える関数
export function toggleActionButtons() {
    const selectedEmployee = document.querySelector("input[name='selectedEmployee']:checked");
    const updateButton = document.getElementById("update-button");
    const deleteButton = document.getElementById("delete-button");
  
    if (selectedEmployee) {
      updateButton.disabled = false;
      deleteButton.disabled = false;
      updateButton.classList.remove("disabled-button");
      deleteButton.classList.remove("disabled-button");
    } else {
      updateButton.disabled = true;
      deleteButton.disabled = true;
      updateButton.classList.add("disabled-button");
      deleteButton.classList.add("disabled-button");
    }
}
  
// ドロップダウンメニューの初期化
export function populateDropdowns(selectedPosition = "", selectedDepartment = "") {
    const positionSelect = document.getElementById("position");
    const departmentSelect = document.getElementById("department");
  
    positionSelect.innerHTML = positions
      .map(pos => `<option value="${pos}" ${pos === selectedPosition ? "selected" : ""}>${pos}</option>`)
      .join("");
  
    departmentSelect.innerHTML = departments
      .map(dep => `<option value="${dep}" ${dep === selectedDepartment ? "selected" : ""}>${dep}</option>`)
      .join("");
}

// 名前の長さをチェック
export function validateNameLength(name, fieldName) {
    if (name.length > NAME_MAX_LENGTH) {
      alert(`${fieldName} は${NAME_MAX_LENGTH}文字以内で入力してください。`);
      return false;
    }
    return true;
}
  
// 役職が有効かチェック
export function validatePosition(position) {
    if (!positions.includes(position)) {
      alert("選択された役職が無効です。");
      return false;
    }
    return true;
}
  
// 部署が有効かチェック
export function validateDepartment(department) {
    if (!departments.includes(department)) {
      alert("選択された部署が無効です。");
      return false;
    }
    return true;
}
  