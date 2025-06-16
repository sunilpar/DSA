package leetcode

import "fmt"

func SearchInsert(nums []int, target int) int {
	l := 0
	h := len(nums) - 1
	mid := l + (h-l)/2
	for l <= h {
		if nums[mid] == target {
			fmt.Printf("mid\n")
			return mid
		}
		if nums[mid] < target {
			l = mid + 1
		} else {
			h = mid - 1
		}
		mid = l + (h-l)/2
	}
	return mid
}
