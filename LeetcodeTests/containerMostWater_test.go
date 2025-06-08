package Leetcodo_Test

import (
	leetcode "DSA/Leetcode/arrays"
	"reflect"
	"testing"
)

func TestMaxArea(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},

		{[]int{1, 1}, 1},

		{[]int{1, 6, 7, 7, 4, 3, 7, 2, 6}, 42},
	}

	for _, test := range tests {
		result := leetcode.MaxArea(test.nums)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input nums=%v,\t expected %v, got %v",
				test.nums, test.expected, result)
		}
	}
}
