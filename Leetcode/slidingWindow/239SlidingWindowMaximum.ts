function maxSlidingWindow(nums: number[], k: number): number[] {
  let expected: number[] = [];
  let max = nums[0];
  for (let left = 0; left <= nums.length - k; left++) {
    let i = left,
      j = left + (k - 1);
    console.log(left, "--------");

    while (i <= j) {
      if (nums[i] > max) {
        console.log(nums[i], ">", max);
        max = nums[i];
      }
      if (nums[j] > max) {
        console.log(nums[j], ">", max);
        max = nums[j];
      }
      i++;
      j--;
    }
    console.log("i=", nums[i], "j=", nums[j], "max=", max);
    expected.push(max);
    max = nums[left + 1];
  }
  return expected;
}

function maxSlidingWindowOptimal(nums: number[], k: number): number[] {
  const result: number[] = [];

  const deque: number[] = [];
  let head = 0;

  for (let i = 0; i < nums.length; i += 1) {
    const windowLeft = i - k + 1;
    if (head < deque.length && deque[head] < windowLeft) {
      head += 1;
    }

    while (head < deque.length && nums[deque[deque.length - 1]] <= nums[i]) {
      deque.pop();
    }

    deque.push(i);

    if (i >= k - 1) {
      result.push(nums[deque[head]]);
    }

    if (head > k) {
      deque.splice(0, head);
      head = 0;
    }
  }

  return result;
}

const nums1 = [1, 3, -1, -3, 5, 3, 6, 7];
const kk1 = 1;
console.log(maxSlidingWindow(nums1, kk1));

// type SlidingWindowTestCase = {
//   nums: number[];
//   k: number;
//   expected: number[];
// };

// const slidingWindowtestCase: SlidingWindowTestCase[] = [
//   // 1. Standard example
//   {
//     nums: [1, 3, -1, -3, 5, 3, 6, 7],
//     k: 3,
//     expected: [3, 3, 5, 5, 6, 7],
//   },
//
//   // 2. Window size is 1
//   {
//     nums: [1, 3, -1, -3, 5, 3, 6, 7],
//     k: 1,
//     expected: [1, 3, -1, -3, 5, 3, 6, 7],
//   },
//
//   // 3. Window size equals array length
//   {
//     nums: [1, 3, -1, -3, 5, 3, 6, 7],
//     k: 8,
//     expected: [7],
//   },
//
//   // 4. Decreasing numbers
//   {
//     nums: [9, 7, 5, 3, 1],
//     k: 2,
//     expected: [9, 7, 5, 3],
//   },
//
//   // 5. Increasing numbers
//   {
//     nums: [1, 2, 3, 4, 5],
//     k: 3,
//     expected: [3, 4, 5],
//   },
//
//   // 6. All elements are equal
//   {
//     nums: [4, 4, 4, 4, 4],
//     k: 2,
//     expected: [4, 4, 4, 4],
//   },
//
//   // 7. Negative numbers
//   {
//     nums: [-7, -8, -3, -5, -2, -6],
//     k: 2,
//     expected: [-7, -3, -3, -2, -2],
//   },
//
//   // 8. Mixed positive and negative numbers
//   {
//     nums: [-1, 5, 2, 7, -3, 4],
//     k: 3,
//     expected: [5, 7, 7, 7],
//   },
//
//   // 9. Maximum appears multiple times
//   {
//     nums: [5, 5, 1, 5, 5],
//     k: 3,
//     expected: [5, 5, 5],
//   },
//
//   // 10. Maximum leaves the window
//   {
//     nums: [10, 1, 2, 3, 4],
//     k: 2,
//     expected: [10, 2, 3, 4],
//   },
//
//   // 11. Alternating high and low values
//   {
//     nums: [1, 10, 1, 10, 1, 10],
//     k: 3,
//     expected: [10, 10, 10, 10],
//   },
//
//   // 12. Single element
//   {
//     nums: [42],
//     k: 1,
//     expected: [42],
//   },
//
//   // 13. Large window with negative values
//   {
//     nums: [-9, -3, -5, -1, -7, -2],
//     k: 4,
//     expected: [-1, -1, -1],
//   },
//
//   // 14. Duplicates at the window boundary
//   {
//     nums: [2, 1, 2, 1, 2, 1],
//     k: 3,
//     expected: [2, 2, 2, 2],
//   },
//
//   // 15. Maximum at the end
//   {
//     nums: [1, 2, 3, 4, 10],
//     k: 2,
//     expected: [2, 3, 4, 10],
//   },
// ];

// let slidingWindowtpassed = 0;
// let slidingWindowtfailed = 0;
//
// for (const [index, test] of slidingWindowtestCase.entries()) {
//   const actual = maxSlidingWindow(test.nums, test.k);
//
//   const isCorrect =
//     actual.length === test.expected.length &&
//     actual.every((value, i) => value === test.expected[i]);
//
//   if (!isCorrect) {
//     slidingWindowtfailed++;
//
//     console.error(`❌ Test #${index + 1} Failed`);
//     console.error(`   nums     : [${test.nums.join(", ")}]`);
//     console.error(`   k        : ${test.k}`);
//     console.error(`   Expected : [${test.expected.join(", ")}]`);
//     console.error(`   Actual   : [${actual.join(", ")}]`);
//     console.error("");
//   } else {
//     slidingWindowtpassed++;
//   }
// }
//
// console.log("================================");
// console.log(`Passed: ${slidingWindowtpassed}`);
// console.log(`Failed: ${slidingWindowtfailed}`);
// console.log(`Total : ${slidingWindowtestCase.length}`);
// console.log("================================");
