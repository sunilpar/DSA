package Algo

import (
	"container/list"
)

type Graph struct {
	vertices int
	adjList  []*list.List
}

func NewGraph(vertices int) *Graph {
	graph := &Graph{
		vertices: vertices,
		adjList:  make([]*list.List, vertices),
	}
	for i := range vertices {
		graph.adjList[i] = list.New()
	}
	return graph
}

func (g *Graph) AddEdge(src, dest int) {
	g.adjList[src].PushBack(dest)
	// Uncomment below for undirected graph
	// g.adjList[dest].PushBack(src)
}

func (g *Graph) DFS(start int, visited []bool, result *[]int) {
	visited[start] = true
	*result = append(*result, start)

	for e := g.adjList[start].Front(); e != nil; e = e.Next() {
		neighbor := e.Value.(int)
		if !visited[neighbor] {
			g.DFS(neighbor, visited, result)
		}
	}
}

func (g *Graph) DFSTraversal(order []int) []int {
	visited := make([]bool, g.vertices)
	result := []int{}

	for _, vertex := range order {
		if !visited[vertex] {
			g.DFS(vertex, visited, &result)
		}
	}
	return result
}
