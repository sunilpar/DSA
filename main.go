package main

import (
	Algo "DSA/Algorithms"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// arr1 := []int{10, 20, 30, 40, 50, 60, 70, 80}
	// arr1 := []int{80, 70, 60, 50, 40, 30, 20, 10}

	fmt.Print("quick sort\n")
	arr1 := rand.Perm(10)
	s := time.Now()
	Algo.QuickSort(arr1, 0, len(arr1)-1)
	//	fmt.Printf("sorted arr %v\n", arr1)
	e := time.Since(s)
	fmt.Printf("time taken:%v\n", e)

	arr1 = rand.Perm(100)
	s = time.Now()
	Algo.QuickSort(arr1, 0, len(arr1)-1)
	//	fmt.Printf("sorted  arr %v\n", arr1)
	e = time.Since(s)
	fmt.Printf("time taken:%v\n", e)

	arr1 = rand.Perm(1000)
	s = time.Now()
	Algo.QuickSort(arr1, 0, len(arr1)-1)
	//	fmt.Printf("sorted  arr %v\n", arr1)
	e = time.Since(s)
	fmt.Printf("time taken:%v\n", e)

	arr1 = rand.Perm(10000)
	s = time.Now()
	Algo.QuickSort(arr1, 0, len(arr1)-1)
	//	fmt.Printf("sorted  arr %v\n", arr1)
	e = time.Since(s)
	fmt.Printf("time taken:%v\n", e)

	fmt.Print("\nmerge sort:\n ")

	arr1 = rand.Perm(10)
	s = time.Now()
	Algo.MergeSort(arr1, 0, len(arr1)-1)
	//	fmt.Printf("sorted arr %v\n", arr1)
	e = time.Since(s)
	fmt.Printf("time taken:%v\n", e)

	arr1 = rand.Perm(100)
	s = time.Now()
	Algo.MergeSort(arr1, 0, len(arr1)-1)
	//	fmt.Printf("sorted  arr %v\n", arr1)
	e = time.Since(s)
	fmt.Printf("time taken:%v\n", e)

	arr1 = rand.Perm(1000)
	s = time.Now()
	Algo.MergeSort(arr1, 0, len(arr1)-1)
	//	fmt.Printf("sorted  arr %v\n", arr1)
	e = time.Since(s)
	fmt.Printf("time taken:%v\n", e)

	arr1 = rand.Perm(10000)
	s = time.Now()
	Algo.MergeSort(arr1, 0, len(arr1)-1)
	//	fmt.Printf("sorted  arr %v\n", arr1)
	e = time.Since(s)
	fmt.Printf("time taken:%v\n", e)

}
