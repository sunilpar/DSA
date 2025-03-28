package Algo

import (
	"fmt"
	"time"

	"golang.org/x/exp/constraints"
)

func BubbleSortFor[T constraints.Ordered](arr []T) ([]T, error) {
	if len(arr) == 0 {
		return nil, fmt.Errorf("cant sort from empty array!!\n")
	}
	s := time.Now()
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	e := time.Since(s)
	fmt.Printf("\ntime taken:%v\n", e)
	return arr, nil
}

func BubbleSort[T constraints.Ordered](arr []T) ([]T, error) {
	if len(arr) == 0 {
		return nil, fmt.Errorf("cant sort from empty array!!\n")
	}
	s := time.Now()
	n := len(arr)

	for i := range arr[:n-1] {
		for j := range arr[:n-i-1] {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	e := time.Since(s)
	fmt.Printf("\ntime taken:%v\n", e)
	return arr, nil
}
