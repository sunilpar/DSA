package DS

import (
	"fmt"
	"log"
)

type Adjlist[T any] struct {
	Nodes []ListNode[T]
}
type ListNode[T any] struct {
	Value  T
	Childs []*ListEdge[T]
}

type ListEdge[T any] struct {
	To     T
	Weight int
}

func CreateAdjList[T comparable]() *Adjlist[T] {
	return &Adjlist[T]{}
}

func (g *Adjlist[T]) rootList(value T) error {
	if len(g.Nodes) != 0 {
		return fmt.Errorf("graph already has other nodes!!")
	}
	r := ListNode[T]{Value: value}
	g.Nodes = append(g.Nodes, r)
	return nil
}

func (g *Adjlist[T]) AddNode(value T) error {
	if len(g.Nodes) == 0 {
		err := g.rootList(value)
		if err != nil {
			log.Fatalf("error: %s\n", err.Error())
		}
	}
	n := ListNode[T]{Value: value}
	g.Nodes = append(g.Nodes, n)
	return nil
}
