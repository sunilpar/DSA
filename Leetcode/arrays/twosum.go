// https://leetcode.com/problems/two-sum/?envType=problem-list-v2&envId=array
package leetcode

func TwoSum(nums []int, target int) []int {
	mymap := make(map[int]int)
	for i, _ := range nums {
		j, ok := mymap[target-nums[i]]
		if ok {
			result := []int{j, i}
			return result
		}
		mymap[nums[i]] = i
	}
	result := []int{-1, -1}
	return result
}
