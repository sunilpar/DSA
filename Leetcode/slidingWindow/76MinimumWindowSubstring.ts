function minWindow(s: string, t: string): string {
  return "";
}

type TestCase = {
  s: string;
  t: string;
  expected: string;
};

const testCases: TestCase[] = [
  {
    s: "ADOBECODEBANC",
    t: "ABC",
    expected: "BANC",
  },
  {
    s: "a",
    t: "a",
    expected: "a",
  },
  {
    s: "a",
    t: "aa",
    expected: "",
  },
  {
    s: "ab",
    t: "b",
    expected: "b",
  },
  {
    s: "aa",
    t: "aa",
    expected: "aa",
  },
  {
    s: "abc",
    t: "ac",
    expected: "abc",
  },
  {
    s: "bba",
    t: "ab",
    expected: "ba",
  },
  {
    s: "cabwefgewcwaefgcf",
    t: "cae",
    expected: "cwae",
  },
  {
    s: "aaaaaaaaaaaabbbbbcdd",
    t: "abcdd",
    expected: "abbbbbcdd",
  },
  {
    s: "abc",
    t: "d",
    expected: "",
  },
  {
    s: "AA",
    t: "AA",
    expected: "AA",
  },
  {
    s: "ADOBECODEBANC",
    t: "AABC",
    expected: "ADOBECODEBA",
  },
  {
    s: "aaflslflsldkalskaaa",
    t: "aaa",
    expected: "aaa",
  },
  {
    s: "abbbbbbbc",
    t: "abc",
    expected: "abbbbbbbc",
  },
  {
    s: "xyyzyzyx",
    t: "xyz",
    expected: "zyx",
  },
];

let passed = 0;
let failed = 0;

for (const [index, test] of testCases.entries()) {
  const actual = minWindow(test.s, test.t);

  if (actual !== test.expected) {
    failed++;

    console.error(`❌ Test #${index + 1} Failed`);
    console.error(`   s        : "${test.s}"`);
    console.error(`   t        : "${test.t}"`);
    console.error(`   Expected : "${test.expected}"`);
    console.error(`   Actual   : "${actual}"`);
    console.error("");
  } else {
    passed++;

    console.log(`✅ Test #${index + 1} Passed`);
  }
}

console.log("");
console.log("================================");
console.log(`Passed: ${passed}`);
console.log(`Failed: ${failed}`);
console.log(`Total : ${testCases.length}`);
console.log("================================");
