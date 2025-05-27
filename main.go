package main

import (
	leetcode "DSA/Leetcode/arrays"
	"fmt"
)

func main() {
	nums1 := []int{1, 2, 4}
	nums2 := []int{3}
	m := leetcode.FindMedianSortedArrays(nums1, nums2)
	fmt.Printf("median:%v\n", m)
}
