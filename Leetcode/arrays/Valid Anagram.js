var s = "conversation";
var t = "conversatin";
// function sortString(str: string): string {
//   return str
//     .split("")
//     .sort((a, b) => a.charCodeAt(0) - b.charCodeAt(0))
//     .join("");
// }
//
// function isAnagram(s: string, t: string): boolean {
//   const sortedS = sortString(s);
//   const sortedT = sortString(t);
//   if (sortedS == sortedT) {
//     return true;
//   } else {
//     return false;
//   }
// }
//sort approach
//
function isAnagram(s, t) {
    if (s.length !== t.length)
        return false;
    var count = {};
    for (var _i = 0, s_1 = s; _i < s_1.length; _i++) {
        var char = s_1[_i];
        count[char] = (count[char] || 0) + 1;
    }
    for (var _a = 0, t_1 = t; _a < t_1.length; _a++) {
        var char = t_1[_a];
        if (!count[char])
            return false;
        count[char]--;
    }
    return true;
}
var r = isAnagram(s, t);
console.log(r);
