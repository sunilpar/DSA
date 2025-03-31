package Algo

import (
	"golang.org/x/exp/constraints"
)

// Merge function to merge two sorted subarrays
func merge[T constraints.Ordered](arr []T, left, mid, right int) {
	leftSize := mid - left + 1
	rightSize := right - mid

	// Create temp slices
	L := make([]T, leftSize)
	R := make([]T, rightSize)

	// Copy data
	copy(L, arr[left:mid+1])
	copy(R, arr[mid+1:right+1])

	i, j, k := 0, 0, left
	for i < leftSize && j < rightSize {
		if L[i] <= R[j] {
			arr[k] = L[i]
			i++
		} else {
			arr[k] = R[j]
			j++
		}
		k++
	}

	// Copy remaining elements
	for i < leftSize {
		arr[k] = L[i]
		i++
		k++
	}
	for j < rightSize {
		arr[k] = R[j]
		j++
		k++
	}
}

// MergeSort function
func MergeSort[T constraints.Ordered](arr []T, left, right int) {
	if left < right {
		mid := left + (right-left)/2
		MergeSort(arr, left, mid)
		MergeSort(arr, mid+1, right)
		merge(arr, left, mid, right)
	}
}
