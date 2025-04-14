package Algo

import (
	DS "DSA/DataStructures"
	"fmt"
	"log"
)

type stack[T comparable] struct {
	Items DS.Stack[T]
}

func NewStack[T comparable]() *stack[T] {
	return &stack[T]{Items: DS.Stack[T]{}}
}

// NOTE: create this type of funciton for all your datatypes for better intalization and not wonder how to make a data type
func (stk *stack[T]) Push(val T) {
	stk.Items.Push(val)
}

func (stk *stack[T]) Pop() (T, error) {
	return stk.Items.Pop()
}
func DFS[T comparable](graph *DS.Graph[T], key T) (*DS.GraphNode[T], []*DS.GraphNode[T], error) {
	if graph.Nodes.Size() == 0 {
		return nil, nil, fmt.Errorf("Graph is empty:%v\n", graph)
	}
	head := graph.Nodes.Head.Value
	s := NewStack[*DS.GraphNode[T]]()
	path := []*DS.GraphNode[T]{}
	s.Items.Push(head)
	for {
		n, err := s.Items.Pop()
		if err != nil {
			log.Fatalf("couldn't pop from the stack: %s\n", err.Error())
		}
		path = append(path, n)
		if n.Value == key {
			return n, path, nil
		}
		for _, c := range n.Childs {
			s.Items.Push(c)
		}
		if s.Items.Size() == 0 {
			break
		}
	}

	return nil, nil, fmt.Errorf("coudn't find %+v\n", key)
}

//pop
