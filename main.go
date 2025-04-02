package main

import (
	Algo "DSA/Algorithms"
	"fmt"
	"time"
)

func main() {
	// arr1 := []int{10, 20, 30, 40, 50, 60, 70, 80}
	// arr1 := []int{80, 70, 60, 50, 40, 30, 20, 10}
	// arr1 := rand.Perm(10)

	arr1 := make([]int, 10)
	for i := range arr1 {
		arr1[i] = i
	}

	fmt.Print("binary search ")
	s := time.Now()
	v := Algo.Binarysearch(arr1, 0, len(arr1)-1, 9)
	if v == -1 {
		fmt.Print("no element found\n")
	} else {
		fmt.Printf("element found at:arr[%v]:%v\n", v, arr1[v])
	}
	e := time.Since(s)
	fmt.Printf("time taken:%v\n", e)

	arr1 = make([]int, 100)
	for i := range arr1 {
		arr1[i] = i
	}

	s = time.Now()
	v = Algo.Binarysearch(arr1, 0, len(arr1)-1, 90)
	if v == -1 {
		fmt.Print("no element found\n")
	} else {
		fmt.Printf("element found at:arr[%v]:%v\n", v, arr1[v])
	}
	e = time.Since(s)
	fmt.Printf("time taken:%v\n", e)

	arr1 = make([]int, 1000)
	for i := range arr1 {
		arr1[i] = i
	}

	s = time.Now()
	v = Algo.Binarysearch(arr1, 0, len(arr1)-1, 900)
	if v == -1 {
		fmt.Print("no element found\n")
	} else {
		fmt.Printf("element found at:arr[%v]:%v\n", v, arr1[v])
	}
	e = time.Since(s)
	fmt.Printf("time taken:%v\n", e)

	arr1 = make([]int, 10000)
	for i := range arr1 {
		arr1[i] = i
	}

	s = time.Now()
	v = Algo.Binarysearch(arr1, 0, len(arr1)-1, 9000)
	if v == -1 {
		fmt.Print("no element found\n")
	} else {
		fmt.Printf("element found at:arr[%v]:%v\n", v, arr1[v])
	}
	e = time.Since(s)
	fmt.Printf("time taken:%v\n", e)
}
