package main

import (
	leetcode "DSA/Leetcode/arrays"
	"fmt"
)

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	// nums := []int{1, 1, 2}
	u := leetcode.RemoveDuplicates(nums)
	fmt.Printf("unq:%v\n", u)

}
