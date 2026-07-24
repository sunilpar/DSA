const nums = [1, 2, 3, 4];
//Output: [24,12,8,6]
const nums1 = [-1, 1, 0, -3, 3];
//Output: [0,0,9,0,0]
function productExceptSelf(nums) {
    const len = nums.length;
    const res = new Array(len);
    let leftProduct = 1;
    for (let i = 0; i < len; i++) {
        res[i] = leftProduct;
        leftProduct *= nums[i];
        console.log("res-- ", res);
    }
    console.log("left product == ", res);
    let rightProduct = 1;
    for (let i = len - 1; i >= 0; i--) {
        res[i] *= rightProduct;
        rightProduct *= nums[i];
        console.log("res-- ", res);
    }
    console.log("right product == ", res);
    return res;
}
console.log(productExceptSelf(nums));
