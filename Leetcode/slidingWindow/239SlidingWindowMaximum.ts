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
