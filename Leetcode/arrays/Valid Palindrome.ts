const str = "A man, a plan, a canal: Panama";
const str1 = "race a car";
const str2 = " ";
const str3 = "0P";

function isPalindrome(s: string): boolean {
  const formated = s.replace(/[^a-zA-Z\d]/g, "").toLowerCase();
  let res = false;

  if (formated.length > 0) {
    let i = 0;
    let j = formated.length - 1;
    while (i <= j) {
      if (formated[i] == formated[j]) {
        res = true;
      } else {
        res = false;
        return res;
      }
      i++;
      j--;
    }
  } else {
    res = true;
  }
  return res;
}

console.log(isPalindrome(str3));
