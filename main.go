package main

import (
	Algo "DSA/Algorithms"
	"fmt"
	"time"
)

func main() {
	arr1 := []int{10, 20, 30, 40, 50, 60, 70, 80}
	// arr1 := []int{80, 70, 60, 50, 40, 30, 20, 10}
	// arr1 := rand.Perm(10)

	fmt.Print("binary search ")
	s := time.Now()
	v := Algo.Binarysearch(arr1, 0, len(arr1)-1, 10)
	if v == -1 {
		fmt.Print("no element found\n")
	} else {
		fmt.Printf("element found at:arr[%v]:%v\n", v, arr1[v])
	}
	e := time.Since(s)
	fmt.Printf("time taken:%v\n", e)
}
