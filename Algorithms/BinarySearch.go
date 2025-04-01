package Algo

import (
	"golang.org/x/exp/constraints"
)

func Binarysearch[T constraints.Ordered](arr []T, low, high int, key T) int {
	for low <= high {
		mid := low + (high-low)/2

		if arr[mid] == key {
			return mid
		}
		if arr[mid] < key {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
