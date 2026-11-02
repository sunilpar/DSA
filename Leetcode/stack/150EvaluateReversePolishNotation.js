function evalRPN(tokens) {
    const stack = [];
    const op = ["+", "-", "*", "/"];
    for (const s of tokens) {
        if (op.includes(s)) {
            const b = Number(stack.pop());
            const a = Number(stack.pop());
            let res = 0;
            switch (s) {
                case "+":
                    res = a + b;
                    break;
                case "-":
                    res = a - b;
                    break;
                case "*":
                    res = a * b;
                    break;
                case "/":
                    res = Math.trunc(a / b);
                    break;
                default:
                    break;
            }
            stack.push(res);
        }
        else {
            stack.push(Number(s));
        }
    }
    return stack.pop() || 0;
}
// Test runner
function runTest(tokens, expected) {
    const result = evalRPN(tokens);
    const passed = result === expected;
    console.log(passed ? "✅ PASSED" : "❌ FAILED");
    console.log("Input:    ", tokens);
    console.log("Expected: ", expected);
    console.log("Received: ", result);
    console.log("------------------------------");
}
// TEST 1: Example 1 — basic multiplication and addition
runTest(["2", "1", "+", "3", "*"], 9);
// TEST 2: Example 2 — division and addition
runTest(["4", "13", "5", "/", "+"], 6);
// TEST 3: Example 3 — complex expression
runTest(["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"], 22);
// TEST 4: Single number
runTest(["42"], 42);
// TEST 5: Negative numbers
runTest(["-2", "3", "*"], -6);
// TEST 6: Subtraction order matters
runTest(["10", "3", "-"], 7);
// TEST 7: Division truncates toward zero
runTest(["7", "2", "/"], 3);
// TEST 8: Negative division truncates toward zero
runTest(["-7", "2", "/"], -3);
// TEST 9: Negative divisor
runTest(["7", "-2", "/"], -3);
// TEST 10: Both operands negative
runTest(["-7", "-2", "/"], 3);
// TEST 11: Multiple operations
runTest(["5", "1", "2", "+", "4", "*", "+", "3", "-"], 14);
// TEST 12: Division before subtraction
runTest(["15", "7", "1", "1", "+", "-", "/", "3", "*", "2", "1", "1", "+", "+", "-"], 5);
// TEST 13: Zero as an operand
runTest(["0", "3", "+"], 3);
// TEST 14: Negative intermediate result
runTest(["3", "5", "-", "2", "*"], -4);
// TEST 15: Larger expression with repeated division
runTest(["100", "10", "/", "5", "/"], 2);
