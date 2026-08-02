const s = "abccdefghijjk";
const s1 = "dvdf";
const s2 = "pwwkew";
const s3 = "abcabcbb";
function lengthOfLongestSubstring(s) {
    const set = new Set();
    let left = 0;
    let longest = 0;
    for (let right = 0; right < s.length; right++) {
        while (set.has(s[right])) {
            set.delete(s[left]);
            left++;
        }
        set.add(s[right]);
        longest = Math.max(longest, right - left + 1);
    }
    return longest;
}
console.log(lengthOfLongestSubstring(s3));
