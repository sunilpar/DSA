// https://leetcode.com/problems/two-sum/?envType=problem-list-v2&envId=array
package leetcode

import "fmt"

func TwoSum(nums []int, target int) []int {
	for i, _ := range nums {
		diff := target - nums[i]
		for j, _ := range nums {
			if nums[j] == diff {
				fmt.Printf("tar:%v found:[%v,%v]\\n", target, i, j)
				return []int{i, j}
			}
		}
	}
	return []int{}

}
