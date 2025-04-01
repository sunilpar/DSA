package Algo

import (
	"golang.org/x/exp/constraints"
)

func merge[T constraints.Ordered](arr []T, left, mid, right int) {
	leftSize := mid - left + 1
	rightSize := right - mid

	L := make([]T, leftSize)
	R := make([]T, rightSize)

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

func MergeSort[T constraints.Ordered](arr []T, left, right int) {
	if left < right {
		mid := left + (right-left)/2
		MergeSort(arr, left, mid)
		MergeSort(arr, mid+1, right)
		merge(arr, left, mid, right)
	}
}
