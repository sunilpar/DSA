const p = [7, 1, 5, 3, 6, 4];
const p1 = [10, 8, 6, 2, 20];
const p2 = [3, 8, 1, 10, 2, 15];
const p3 = [5, 5, 5, 5];
const p4 = [5];
const p5 = [7, 6, 4, 3, 1];
const p6 = [1, 4, 2];
const p7 = [2, 1, 2, 1, 0, 1, 2];

function maxProfit(p: number[]): number {
  let l = 0,
    r = 1,
    maxp = 0;

  while (r < p.length) {
    if (p[r] - p[l] > maxp) {
      maxp = p[r] - p[l];
    }

    if (p[l] > p[r]) {
      l = r;
      r++;
    } else {
      r++;
    }
  }
  return maxp;
}
console.log(maxProfit(p7));
