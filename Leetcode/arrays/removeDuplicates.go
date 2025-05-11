package leetcode

func RemoveDuplicates(nums []int) int {
	j := 0
	for _, v := range nums {
		if nums[j] < v {
			j++
		}
		nums[j], v = v, nums[j]
	}
	return j + 1
}

//https://leetcode.com/problems/remove-duplicates-from-sorted-array/?envType=problem-list-v2&envId=array
