const data = [1, 8, 6, 2, 5, 4, 8, 3, 7];
const data1 = [1, 1];
const data2 = [4, 3, 2, 1, 4];
const data3 = [1, 2, 1];

function maxArea(height: number[]): number {
  let left = 0;
  let right = height.length - 1;

  let maxWater = 0;

  while (left < right) {
    const area = Math.min(height[left], height[right]) * (right - left);
    maxWater = Math.max(maxWater, area);

    if (height[left] < height[right]) {
      left++;
    } else {
      right--;
    }
  }

  return maxWater;
}

console.log(maxArea(data));
