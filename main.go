package main

import (
	Algo "DSA/Algorithms"
	"fmt"
)

func main() {
	maze := [][]string{
		{"x", "x", "x", "x", "x", "x", "x", "x", "x", "x", " ", "x"},
		{"x", " ", " ", " ", " ", " ", " ", " ", " ", "x", " ", "x"},
		{"x", " ", " ", " ", " ", " ", " ", " ", " ", "x", " ", "x"},
		{"x", " ", "x", "x", "x", "x", "x", "x", "x", "x", " ", "x"},
		{"x", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", "x"},
		{"x", " ", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x"},
	}

	path := Algo.SolveMaze(maze, "x", []int{0, 10}, []int{5, 1})

	fmt.Println(path)

}
