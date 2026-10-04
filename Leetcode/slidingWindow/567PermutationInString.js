function checkinclusion(s1, s2) {
    if (s1.length > s2.length) {
        return false;
    }
    const s2map = new Map();
    s2.split("").forEach((char) => {
        var _a;
        s2map.set(char, ((_a = s2map.get(char)) !== null && _a !== void 0 ? _a : 0) + 1);
    });
    s1.split("").forEach((char) => {
        const count = s2map.get(char);
        if (!count || count < 1) {
            return false;
        }
        else {
            s2map.set(char, count - 1);
        }
    });
    return true;
}
const testCases = [
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
    const actual = checkinclusion(test.s1, test.s2);
    if (actual !== test.expected) {
        failed++;
        console.error(`❌ Test #${index + 1} Failed`);
        console.error(`   s1       : "${test.s1}"`);
        console.error(`   s2       : "${test.s2}"`);
        console.error(`   Expected : ${test.expected}`);
        console.error(`   Actual   : ${actual}`);
        console.error("");
    }
    else {
        passed++;
    }
}
console.log("================================");
console.log(`Passed: ${passed}`);
console.log(`Failed: ${failed}`);
console.log(`Total : ${testCases.length}`);
console.log("================================");
