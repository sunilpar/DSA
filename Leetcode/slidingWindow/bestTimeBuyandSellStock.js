const p = [7, 1, 5, 3, 6, 4];
const p1 = [10, 8, 6, 2, 20];
const p2 = [3, 8, 1, 10, 2, 15];
const p3 = [5, 5, 5, 5];
const p4 = [5];
const p5 = [7, 6, 4, 3, 1];
const p6 = [1, 4, 2];
const p7 = [2, 1, 2, 1, 0, 1, 2];
// function maxProfit(p: number[]): number {
//   let i = 0;
//   let j = p.length - 1;
//   let mini = 0;
//   let maxi = p.length - 1;
//   //   finding max and min
//   while (i < j) {
//     if (p[i] < p[mini]) {
//       mini = i;
//     }
//
//     if (p[j] > p[maxi]) {
//       maxi = j;
//     }
//
//     //updating i and j
//     if (p[i] < p[j]) {
//       i++;
//     } else {
//       i++;
//       j--;
//     }
//   }
//
//   if (p[maxi] - p[mini] < 0) {
//     return 0;
//   } else {
//     return p[maxi] - p[mini];
//   }
// }
function maxProfit(p) {
    let l = 0, r = 1, maxp = 0;
    while (r < p.length) {
        if (p[r] - p[l] > maxp) {
            maxp = p[r] - p[l];
        }
        if (p[l] > p[r]) {
            l = r;
            r++;
        }
        else {
            r++;
        }
    }
    return maxp;
}
console.log(maxProfit(p7));
