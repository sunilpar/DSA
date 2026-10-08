function minWindow(s, t) {
    var _a, _b, _c, _d, _e;
    if (t.length > s.length) {
        return "";
    }
    const need = new Map();
    const window = new Map();
    for (const char of t) {
        need.set(char, ((_a = need.get(char)) !== null && _a !== void 0 ? _a : 0) + 1);
    }
    let left = 0;
    let formed = 0;
    const required = need.size;
    let minLength = Infinity;
    let minStart = 0;
    for (let right = 0; right < s.length; right++) {
        const char = s[right];
        window.set(char, ((_b = window.get(char)) !== null && _b !== void 0 ? _b : 0) + 1);
        if (need.has(char) && window.get(char) === need.get(char)) {
            formed++;
        }
        while (formed === required) {
            const currentLength = right - left + 1;
            if (currentLength < minLength) {
                minLength = currentLength;
                minStart = left;
            }
            const leftChar = s[left];
            window.set(leftChar, ((_c = window.get(leftChar)) !== null && _c !== void 0 ? _c : 0) - 1);
            if (need.has(leftChar) &&
                ((_d = window.get(leftChar)) !== null && _d !== void 0 ? _d : 0) < ((_e = need.get(leftChar)) !== null && _e !== void 0 ? _e : 0)) {
                formed--;
            }
            left++;
        }
    }
    if (minLength === Infinity) {
        return "";
    }
    return s.slice(minStart, minStart + minLength);
}
const testCases = [
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
    }
    else {
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
