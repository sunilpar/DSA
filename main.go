package main

import (
	Algo "DSA/Algorithms"
	"fmt"
	"time"
)

func main() {
	//default is assending order
	arr1 := []int{10, 20, 30, 40, 50, 60, 70, 80}
	// arr1 := []int{80, 70, 60, 50, 40, 30, 20, 10}

	s := time.Now()
	Algo.QuickSort(arr1, 0, len(arr1)-1)
	//	fmt.Printf("sorted arr %v\n", arr1)
	e := time.Since(s)
	fmt.Printf("time taken:%v\n", e)
}
