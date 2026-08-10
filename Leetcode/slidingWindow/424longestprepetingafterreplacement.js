const tests = [
    ["A", 0],
    ["A", 100],
    ["AAAAAA", 0],
    ["AAAAAA", 3],
    ["AB", 1],
    ["AB", 0],
    ["ABAB", 2],
    ["AABABBA", 1],
    ["ABCDE", 0],
    ["ABCDE", 1],
    ["ABCDE", 4],
    ["ABABABAB", 1],
    ["ABABABAB", 2],
    ["ABABABAB", 3],
    ["ABCDEFG", 2],
    ["BAAAAB", 1],
    ["BAAAAB", 2],
    ["AAAAABAAAA", 1],
    ["AAAAABAAAA", 0],
    ["ABCDDDDDD", 2],
    ["DDDDDDABC", 2],
    ["ABBB", 2],
    ["ABBB", 1],
    ["ABCDEFF", 1],
    ["AAABBC", 2],
    ["AABBBCC", 2],
    ["AABCAA", 1],
    ["BBBBB", 100],
    ["ZZZZYZZZZ", 1],
    ["BAAAB", 2],
];
const str = "AAAA";
const k = 2;
const str1 = "AAABBC";
const k1 = 2;
const str2 = "ABCDDDDDD";
const k2 = 2;
function characterReplacementBRUT(s, k) {
    if (k > s.length) {
        return s.length;
    }
    let max = 0;
    let i = 0;
    let j = 1;
    let tempmax = 1;
    let tempk = k;
    while (i < s.length - 1) {
        if (j > s.length - 1) {
            tempmax += tempk;
            i++;
            j = i + 1;
            if (tempmax > max) {
                max = tempmax;
            }
            tempmax = 1;
            tempk = k;
            continue;
        }
        if (s[i] == s[j]) {
            tempmax++;
            j++;
        }
        else {
            if (tempk != 0) {
                tempk--;
                tempmax++;
                j++;
            }
            else {
                i++;
                j = i + 1;
                tempmax = 1;
                tempk = k;
                continue;
            }
        }
        if (tempmax > max) {
            max = tempmax;
        }
    }
    return max;
}
function characterReplacement(s, k) {
    var _a;
    const count = new Map();
    let left = 0;
    let maxFreq = 0;
    let longest = 0;
    for (let right = 0; right < s.length; right++) {
        const ch = s[right];
        count.set(ch, ((_a = count.get(ch)) !== null && _a !== void 0 ? _a : 0) + 1);
        maxFreq = Math.max(maxFreq, count.get(ch));
        while (right - left + 1 - maxFreq > k) {
            const leftChar = s[left];
            count.set(leftChar, count.get(leftChar) - 1);
            left++;
        }
        longest = Math.max(longest, right - left + 1);
    }
    return longest;
}
console.log(characterReplacement(str, k));
