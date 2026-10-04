function checkInclusion(s1: string, s2: string): boolean {
  if (s1.length > s2.length) {
    return false;
  }

  const s1Count = new Array(26).fill(0);
  const windowCount = new Array(26).fill(0);

  // Count characters in s1
  for (const char of s1) {
    s1Count[char.charCodeAt(0) - 97]++;
  }

  // Create the first window in s2
  for (let i = 0; i < s1.length; i++) {
    windowCount[s2.charCodeAt(i) - 97]++;
  }

  // Check first window
  if (s1Count.toString() === windowCount.toString()) {
    return true;
  }

  // Slide the window
  for (let i = s1.length; i < s2.length; i++) {
    // Add new character
    windowCount[s2.charCodeAt(i) - 97]++;

    // Remove character leaving the window
    windowCount[s2.charCodeAt(i - s1.length) - 97]--;

    // Compare frequencies
    if (s1Count.toString() === windowCount.toString()) {
      return true;
    }
  }

  return false;
}
type TestCase = {
  s1: string;
  s2: string;
  expected: boolean;
};

const testCases: TestCase[] = [
  {
    s1: "ab",
    s2: "eidbaooo",
    expected: true,
  },
  {
    s1: "ab",
    s2: "eidboaoo",
    expected: false,
  },
  {
    s1: "adc",
    s2: "dcda",
    expected: true,
  },
  {
    s1: "hello",
    s2: "ooolleoooleh",
    expected: false,
  },
  {
    s1: "xyz",
    s2: "afdgzyxksldfm",
    expected: true,
  },
  {
    s1: "a",
    s2: "b",
    expected: false,
  },
  {
    s1: "abc",
    s2: "ccccbbbbaaaa",
    expected: false,
  },
  {
    s1: "abc",
    s2: "bbbca",
    expected: true,
  },
  {
    s1: "abcd",
    s2: "dcba",
    expected: true,
  },
  {
    s1: "aaa",
    s2: "aaabaaaa",
    expected: true,
  },
  {
    s1: "abc",
    s2: "cbaebabacd",
    expected: true,
  },
  {
    s1: "abcd",
    s2: "abc",
    expected: false,
  },
  {
    s1: "a",
    s2: "a",
    expected: true,
  },
  {
    s1: "aa",
    s2: "ab",
    expected: false,
  },
  {
    s1: "aa",
    s2: "baa",
    expected: true,
  },
  {
    s1: "abc",
    s2: "defghijkl",
    expected: false,
  },
  {
    s1: "abc",
    s2: "cab",
    expected: true,
  },
  {
    s1: "zzzz",
    s2: "zzzzzz",
    expected: true,
  },
];

let passed = 0;
let failed = 0;

for (const [index, test] of testCases.entries()) {
  const actual = checkInclusion(test.s1, test.s2);

  if (actual !== test.expected) {
    failed++;

    console.error(`❌ Test #${index + 1} Failed`);
    console.error(`   s1       : "${test.s1}"`);
    console.error(`   s2       : "${test.s2}"`);
    console.error(`   Expected : ${test.expected}`);
    console.error(`   Actual   : ${actual}`);
    console.error("");
  } else {
    passed++;
  }
}

console.log("================================");
console.log(`Passed: ${passed}`);
console.log(`Failed: ${failed}`);
console.log(`Total : ${testCases.length}`);
console.log("================================");
