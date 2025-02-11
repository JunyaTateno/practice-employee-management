import { setupEmployeePage } from "./employee.js";
import { setupRegisterForm } from "./register.js";
import { loadUpdateForm, setupUpdateForm } from "./update.js";
import { toggleActionButtons } from "./utils.js";

// ページ読み込み時の処理
document.addEventListener("DOMContentLoaded", () => {
  const pageId = document.body.id;

  if (pageId === "employee-management") {
    setupEmployeePage();
    toggleActionButtons();
  } else if (pageId === "register-page") {
    setupRegisterForm();
  } else if (pageId === "update-page") {
    loadUpdateForm();
    setupUpdateForm();
  }

});

// ラジオボタンの変更時に `toggleActionButtons()` を実行
document.addEventListener("change", (event) => {
  if (event.target.name === "selectedEmployee") {
    toggleActionButtons();
  }
});
