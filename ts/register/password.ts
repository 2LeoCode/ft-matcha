import { narrow } from "../utils";

type PasswordRequirements = Readonly<{
  [Req in [
    "min-length",
    "max-length",
    "lower",
    "upper",
    "digit",
    "spec",
  ][number]]: boolean;
}>;

function validatePassword(value: string): PasswordRequirements {
  return {
    "min-length": value.length >= 8,
    "max-length": value.length <= 32,
    "lower": /[a-z]/.test(value),
    "upper": /[A-Z]/.test(value),
    "digit": /[0-9]/.test(value),
    "spec": /[^a-zA-Z0-9]/.test(value),
  };
}

const $password = document.getElementById(
  "register:password",
)! as HTMLInputElement;

export let passwordRequirements: PasswordRequirements = validatePassword(
  $password.value,
);

// Live feedback for password validity
$password.addEventListener("input", () => {
  passwordRequirements = validatePassword($password.value);

  for (const requirement in passwordRequirements) {
    narrow<keyof PasswordRequirements>(requirement);
    const $req = document.getElementById(
      `register:password:req:${requirement}`,
    )!;
    if (passwordRequirements[requirement]) $req.style.display = "none";
    else $req.style.display = "block";
  }
});
