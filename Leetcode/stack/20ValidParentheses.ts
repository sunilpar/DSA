const parenthesesMap: Record<string, string> = {
  "(": ")",
  "[": "]",
  "{": "}",
};
function isValid(s: string): boolean {
  const stack: string[] = [];
  for (let i = 0; i < s.length; i++) {
    if (parenthesesMap[stack[stack.length - 1]] !== s[i]) {
      stack.push(s[i]);
    } else {
      stack.pop();
    }
  }
  return stack.length === 0;
}

const s = "()";
const s2 = "()[]{}";
const s3 = "(){{[]{}";

console.log(isValid(s3));
