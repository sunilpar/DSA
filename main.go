package main

import (
	Algo "DSA/Algorithms"
	"fmt"
	"log"
)

func main() {
	//default is assending order
	var arr1 = []int{80, 70, 60, 50, 40, 30, 20, 10}
	var arr2 = []int{10, 20, 30, 40, 50, 60, 70, 80}

	arr, err := Algo.InsertionSort(arr1)
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}
	fmt.Printf("sorted arr %v\n", arr)

	arr, err = Algo.InsertionSort(arr2)
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}
	fmt.Printf("sorted arr %v\n", arr)

}
