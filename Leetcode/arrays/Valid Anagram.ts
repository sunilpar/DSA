const s = "conversation";
const t = "conversatin";

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
function isAnagram(s: string, t: string): boolean {
  if (s.length !== t.length) return false;

  const count: Record<string, number> = {};

  for (const char of s) {
    count[char] = (count[char] || 0) + 1;
  }

  for (const char of t) {
    if (!count[char]) return false;
    count[char]--;
  }

  return true;
}

const r = isAnagram(s, t);
console.log(r);
