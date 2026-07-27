const nums = [0, 3, 7, 2, 5, 8, 4, 6, 0, 1];

const nums1 = [100, 4, 200, 1, 3, 2];

function longestConsecutive(nums: number[]): number {
  const set = new Set(nums);

  let max = 0;

  for (const curr of set) {
    console.log("main");
    if (!set.has(curr - 1)) {
      let next = curr + 1;

      while (set.has(next)) {
        console.log("next ", next);
        next++;
      }

      const sequence = next - curr;

      if (sequence > max) {
        max = sequence;
      }
    }
  }

  return max;
}

console.log(longestConsecutive(nums));
