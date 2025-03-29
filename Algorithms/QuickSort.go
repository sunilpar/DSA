package Algo

import (
	"golang.org/x/exp/constraints"
)

func partition[T constraints.Ordered](arr []T, p, q int) int {
	// r := rand.Intn(q-p+1) + p
	r := p + (q-p)/2
	P := arr[r]
	i := p
	j := q
	for i < j {
		for arr[i] <= P && i <= q-1 {
			i++
		}
		for arr[j] > P && j >= p+1 {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[r], arr[j] = arr[j], arr[r]
	return j
}

func QuickSort[T constraints.Ordered](arr []T, p, q int) {
	if p < q {
		P := partition(arr, p, q)
		QuickSort(arr, p, P-1)
		QuickSort(arr, P+1, q)
	}

}
