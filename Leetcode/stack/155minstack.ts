class MinStack {
  stack: number[];
  mins: number[];

  constructor() {
    this.stack = [];
    this.mins = [];
  }

  push(value: number): void {
    this.stack.push(value);
    const minsLength = this.mins.length;
    if (minsLength === 0 || value < this.mins[minsLength - 1]) {
      this.mins.push(value);
    } else {
      this.mins.push(this.mins[minsLength - 1]);
    }
  }

  pop(): void {
    this.stack.pop();
    this.mins.pop();
  }

  top(): number {
    return this.stack[this.stack.length - 1];
  }

  getMin(): number {
    return this.mins[this.mins.length - 1];
  }
}

// LeetCode-style test runner
type Operation = "MinStack" | "push" | "pop" | "top" | "getMin";

function runTest(
  operations: Operation[],
  args: number[][],
  expected: (number | null)[],
): void {
  const stack = new MinStack();
  const output: (number | null)[] = [];

  for (let i = 0; i < operations.length; i++) {
    const operation = operations[i];
    const operationArgs = args[i];

    switch (operation) {
      case "MinStack":
        output.push(null);
        break;

      case "push":
        stack.push(operationArgs[0]);
        output.push(null);
        break;

      case "pop":
        stack.pop();
        output.push(null);
        break;

      case "top":
        output.push(stack.top());
        break;

      case "getMin":
        output.push(stack.getMin());
        break;
    }
  }

  const passed = JSON.stringify(output) === JSON.stringify(expected);

  console.log(passed ? "✅ PASSED" : "❌ FAILED");
  console.log("Expected:", expected);
  console.log("Received:", output);
  console.log("----------------------------");
}

// TEST CASE 1: LeetCode example
runTest(
  ["MinStack", "push", "push", "push", "getMin", "pop", "top", "getMin"],
  [[], [-2], [0], [-3], [], [], [], []],
  [null, null, null, null, -3, null, 0, -2],
);

// TEST CASE 2: Increasing values
runTest(
  ["MinStack", "push", "push", "push", "top", "getMin", "pop", "getMin"],
  [[], [1], [2], [3], [], [], [], []],
  [null, null, null, null, 3, 1, null, 1],
);

// TEST CASE 3: Your example with multiple new minimums
runTest(
  [
    "MinStack",
    "push",
    "push",
    "push",
    "push",
    "push",
    "push",
    "getMin",
    "pop",
    "getMin",
    "pop",
    "getMin",
    "pop",
    "getMin",
  ],
  [[], [-2], [0], [-3], [1], [4], [-4], [], [], [], [], [], [], []],
  [null, null, null, null, null, null, null, -4, null, -3, null, -3, null, -2],
);

// TEST CASE 4: Duplicate minimums
runTest(
  [
    "MinStack",
    "push",
    "push",
    "push",
    "getMin",
    "pop",
    "getMin",
    "pop",
    "getMin",
  ],
  [[], [2], [2], [1], [], [], [], [], []],
  [null, null, null, null, 1, null, 2, null, 2],
);

// TEST CASE 5: Negative numbers
runTest(
  ["MinStack", "push", "push", "push", "getMin", "pop", "top", "getMin"],
  [[], [-5], [-10], [-3], [], [], [], []],
  [null, null, null, null, -10, null, -10, -10],
);

// TEST CASE 6: Pop until one element remains
runTest(
  ["MinStack", "push", "push", "push", "pop", "pop", "top", "getMin"],
  [[], [8], [3], [6], [], [], [], []],
  [null, null, null, null, null, null, 8, 8],
);

// TEST CASE 7: Large positive and negative values
runTest(
  ["MinStack", "push", "push", "push", "getMin", "pop", "getMin", "top"],
  [[], [2147483647], [-2147483648], [0], [], [], [], []],
  [null, null, null, null, -2147483648, null, -2147483648, -2147483648],
);
