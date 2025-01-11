import { passwordRequirements } from "./password";

const $submit = document.getElementById(
  "register:submit",
)! as HTMLButtonElement;

const $form = document.body.querySelector("form")!;

// Enable submit button only if requirements are met
$form.addEventListener("input", () => {
  if (
    $form.checkValidity() &&
    !Object.values(passwordRequirements).includes(false)
  ) {
    $submit.disabled = false;
  } else {
    $submit.disabled = true;
  }
});
