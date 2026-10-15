const parenthesesMap = {
    "(": ")",
    "[": "]",
    "{": "}",
};
function isValid(s) {
    const stack = [];
    for (let i = 0; i < s.length; i++) {
        if (parenthesesMap[stack[stack.length - 1]] !== s[i]) {
            stack.push(s[i]);
        }
        else {
            stack.pop();
        }
    }
    return stack.length === 0;
}
const s = "()";
const s2 = "()[]{}";
console.log(isValid(s));
