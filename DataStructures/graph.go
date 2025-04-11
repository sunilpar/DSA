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

// func (g *Graph[T]) Insertafter(gn *GraphNode[T], value T) error {
//
//		newNode := &GraphNode[T]{Value: value}
//		fmt.Print("\nnodes graph:")
//		for e := g.Nodes.Front(); e != nil; e = e.Next() {
//			fmt.Printf("%v ", e.Value)
//			if e.Value == gn {
//				g.Nodes.InsertAfter(newNode, e)
//			}
//		}
//		fmt.Printf("\n")
//		gn.Childs = append(gn.Childs, newNode)
//		return nil
//	}
func (g *Graph[T]) Insertafter(gn *GraphNode[T], value T) error {
	newNode := &GraphNode[T]{Value: value}
	fmt.Print("\nGraph Nodes:\n")

	for e := g.Nodes.Front(); e != nil; e = e.Next() {
		prevVal := "nil"
		nextVal := "nil"

		if e.Prev() != nil {
			prevVal = fmt.Sprintf("%v", e.Prev().Value.(*GraphNode[T]).Value)
		}
		if e.Next() != nil {
			nextVal = fmt.Sprintf("%v", e.Next().Value.(*GraphNode[T]).Value)
		}

		currVal := e.Value.(*GraphNode[T]).Value
		fmt.Printf("Prev: %v <- Curr: %v -> Next: %v\n", prevVal, currVal, nextVal)

		if e.Value == gn {
			g.Nodes.InsertAfter(newNode, e)
			fmt.Printf("Inserted new node with value %v after %v\n", newNode.Value, gn.Value)
		}
	}

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
