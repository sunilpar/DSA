let nums = [1, 2, 4, 3, 1];

// function containsDuplicate(nums: number[]): boolean {
//   const temp: number[] = new Array(nums.length);
//
//   for (let n in nums) {
//     for (let j in temp) {
//       if (nums[n] === temp[j]) {
//         return true;
//       }
//     }
//     temp[n] = nums[n];
//   }
//
//   return false;
// }
//
// Brut force

function containsDuplicate(nums: number[]): boolean {
  const seen = new Set<number>();

  for (const num of nums) {
    if (seen.has(num)) {
      return true;
    }
    seen.add(num);
  }

  return false;
}

const bool = containsDuplicate(nums);
console.log(bool);
