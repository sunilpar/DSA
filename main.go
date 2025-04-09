package main

import (
	Algo "DSA/Algorithms"
	"fmt"
)

func main() {
	graph := Algo.NewGraph(4)

	graph.AddEdge(2, 0)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 2)
	graph.AddEdge(0, 1)
	graph.AddEdge(3, 3)
	graph.AddEdge(1, 3)

	order := []int{2, 0, 1, 3}

	fmt.Println("Following is Depth First Traversal (starting from vertex 2):")
	result := graph.DFSTraversal(order)
	for _, v := range result {
		fmt.Printf("%d ", v)
	}
}
