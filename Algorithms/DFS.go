package Algo

import (
	DS "DSA/DataStructures"
)

type stack[T comparable] struct {
	s DS.Stack[T]
}

func NewStack[T comparable]() *stack[T] {
	return &stack[T]{s: DS.Stack[T]{}}
}

// NOTE: create this type of funciton for all your datatypes for better intalization and not wonder how to make a data type
func (stk *stack[T]) Push(val T) {
	stk.s.Push(val)
}

func (stk *stack[T]) Pop() (T, error) {
	return stk.s.Pop()
}

// func DFS[T comparable](graph *DS.Graph[T], key T) (*DS.GraphNode[T], []*DS.GraphNode[T], error) {
// 	if graph.Nodes.Size() == 0 {
// 		return nil, nil, fmt.Errorf("Graph is empty:%v\n", graph)
// 	}
// 	elem := graph.Nodes.Head
// 	head := elem.Value
// 	q := NewStack[T]()
// 	path := []*DS.GraphNode[T]{}
// 	path = append(path, head)
// 	q.enqueue(head)
//
// 	for e := q.list.Front(); e != nil; e = e.Next() {
// 		if e != q.list.Front() {
// 			n, err := q.dequeue()
// 			if err == false {
// 				log.Fatalf("coundn't dequeue!!!\n")
// 			}
// 			path = append(path, n)
// 		}
// 		node := e.Value.(*DS.GraphNode[T])
// 		if node.Value == key {
// 			return node, path, nil
// 		} else {
// 			for _, c := range node.Childs {
// 				q.enqueue(c)
// 			}
// 		}
// 	}
// 	return nil, nil, fmt.Errorf("coudn't find %+v\n", key)
// }
