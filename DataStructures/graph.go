package DS

import (
	"container/list"
	"fmt"
)

type GraphNode[T any] struct {
	Value T

	Childs []*GraphNode[T]
}

type Graph[T any] struct {
	Nodes *list.List
}

func CreateGraph[T any]() *Graph[T] {
	return &Graph[T]{
		Nodes: list.New(),
	}
}

func (g *Graph[T]) Root(value T) *GraphNode[T] {
	node := &GraphNode[T]{Value: value}
	g.Nodes.PushFront(node)
	e := g.Nodes.Front().Value.(*GraphNode[T])
	return e
}

func (g *Graph[T]) InsertAFter(gn *GraphNode[T], value T) error {
	newNode := &GraphNode[T]{Value: value}
	gn.Childs = append(gn.Childs, newNode)
	return nil
}

func (g *Graph[T]) Display() {
	for e := g.Nodes.Front(); e != nil; e = e.Next() {
		node := e.Value.(*GraphNode[T])
		fmt.Printf("%v ", node.Value)
	}
	fmt.Printf("\n")
}
