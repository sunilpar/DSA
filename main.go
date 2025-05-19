package main

import (
	leetcode "DSA/Leetcode/arrays"
)

func main() {
	board := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	leetcode.GameOfLife(board)
}
