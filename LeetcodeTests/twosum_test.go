package Leetcodo_Test

import (
	leetcode "DSA/Leetcode/arrays"
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{1, 2, 3}, 7, nil},
	}

	for _, test := range tests {
		result := leetcode.TwoSum(test.nums, test.target)

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input nums=%v, target=%d; expected %v, got %v",
				test.nums, test.target, test.expected, result)
		}
	}
}
