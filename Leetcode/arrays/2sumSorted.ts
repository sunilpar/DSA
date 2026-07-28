const num = [2, 7, 11, 15];
const t = 9;
const num1 = [2, 3, 4];
const t1 = 6;
const num2 = [-1, 0];
const t2 = -1;

function twoSum(numbers: number[], target: number): number[] {
  let left = 0;
  let right = numbers.length - 1;

  while (left < right) {
    const sum = numbers[left] + numbers[right];

    if (sum === target) {
      return [left + 1, right + 1];
    }

    if (sum < target) {
      left++;
    } else {
      right--;
    }
  }

  return [];
}

console.log(twoSum(num, t));
