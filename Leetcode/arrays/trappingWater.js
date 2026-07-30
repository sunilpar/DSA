const height = [0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1];
const height1 = [4, 2, 0, 3, 2, 5];
//brut
// function trap(height: number[]): number {
//   const maxl: number[] = [];
//   const maxr: number[] = [];
//
//   let water = 0;
//   let j = height.length - 1;
//
//   let mr = 0;
//   let ml = 0;
//   for (let i = 0; i < height.length; i++, j--) {
//     //for adding value
//     if (i == 0 && j == height.length - 1) {
//       maxl.push(0);
//       maxr.push(0);
//     } else {
//       maxl[i] = ml;
//       maxr[j] = mr;
//     }
//     //for updating ml and mr
//     if (ml < height[i]) {
//       ml = height[i];
//     }
//     if (mr < height[j]) {
//       mr = height[j];
//     }
//   }
//   console.log("maxl", maxl);
//   console.log("maxr", maxr);
//
//   for (let i = 0; i < height.length; i++) {
//     let tempw = Math.min(maxl[i], maxr[i]) - height[i];
//     console.log("i-", i, "tempw-", tempw);
//     if (Math.min(maxl[i], maxr[i]) - height[i] > 0) {
//       water += tempw;
//     }
//   }
//
//   return water;
// }
function trap(height) {
    let area = 0;
    let l = 0;
    let maxL = height[l];
    let r = height.length - 1;
    let maxR = height[r];
    while (l < r) {
        if (height[l] > height[r]) {
            r--;
            maxR = Math.max(height[r], maxR);
            area += maxR - height[r];
        }
        else {
            l++;
            maxL = Math.max(height[l], maxL);
            area += maxL - height[l];
        }
    }
    return area;
}
console.log(trap(height));
