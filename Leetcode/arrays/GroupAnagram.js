const words = ["eat", "tea", "tan", "ate", "nat", "bat"];
function groupAnagrams(strs) {
    const groups = new Map();
    for (const word of strs) {
        const freq = new Array(26).fill(0);
        for (const ch of word) {
            freq[ch.charCodeAt(0) - 97]++;
        }
        const key = freq.join("#");
        if (!groups.has(key)) {
            groups.set(key, []);
        }
        groups.get(key).push(word);
    }
    return [...groups.values()];
}
console.log(groupAnagrams(words));
