package Algo

import (
	"fmt"
	"time"

	"golang.org/x/exp/constraints"
)

func InsertionSort[T constraints.Ordered](arr []T) ([]T, error) {
	if len(arr) == 0 {
		return nil, fmt.Errorf("can't sort from empty array!\n")
	}
	s := time.Now()
	for i, key := range arr[1:] {
		j := i
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
	e := time.Since(s)
	fmt.Printf("\nTime taken: %v\n", e)
	return arr, nil
}
