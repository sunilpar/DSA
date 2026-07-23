const nums = [1, 2, 1, 2, 1, 2, 3, 1, 3, 2];
const k = 2;
function topKFrequent(nums, k) {
    var _a;
    const map = new Map();
    for (let i = 0; i < nums.length; i++) {
        map.set(nums[i], ((_a = map.get(nums[i])) !== null && _a !== void 0 ? _a : 0) + 1);
    }
    const sortedMap = new Map([...map.entries()].sort((a, b) => b[1] - a[1]));
    return [...sortedMap.keys()].slice(0, k);
}
console.log(topKFrequent(nums, k));
