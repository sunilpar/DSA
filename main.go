package main

import (
	leetcode "DSA/Leetcode/arrays"
	"fmt"
)

func main() {
	nums := []int{1, 3, 5, 6}
	// target := 5
	// target := 2
	target := 4
	u := leetcode.SearchInsert(nums, target)
	fmt.Printf("main-->\ttarget:%v returned:%v\n", target, u)
}
