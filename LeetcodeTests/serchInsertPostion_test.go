package Leetcodo_Test

import (
	leetcode "DSA/Leetcode/arrays"
	"reflect"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected int
	}{
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
	}

	for _, test := range tests {
		result := leetcode.SearchInsert(test.nums, test.target)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input nums=%v, target=%d;\t expected %v, got %v",
				test.nums, test.target, test.expected, result)
		}
	}
}
