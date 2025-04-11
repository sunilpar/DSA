package Algo

import (
	DS "DSA/DataStructures"
	"container/list"
	"fmt"
	"log"
)

type queue[T comparable] struct {
	list *list.List
}

func newQueue[T comparable]() *queue[T] {
	return &queue[T]{list: list.New()}
}

func (q *queue[T]) enqueue(v *DS.GraphNode[T]) {
	q.list.PushBack(v)
}

func (q *queue[T]) dequeue() (*DS.GraphNode[T], bool) {
	if q.list.Len() == 0 {
		return nil, false
	}
	elem := q.list.Front()
	q.list.Remove(elem)
	return elem.Value.(*DS.GraphNode[T]), true
}

func BFS[T comparable](graph *DS.Graph[T], key T) (*DS.GraphNode[T], []*DS.GraphNode[T], error) {
	if graph.Nodes.Len() == 0 {
		return nil, nil, fmt.Errorf("Graph is empty:%v\n", graph)
	}
	elem := graph.Nodes.Front()
	head := elem.Value.(*DS.GraphNode[T])
	q := newQueue[T]()
	q.enqueue(head)
	path := []*DS.GraphNode[T]{}
	path = append(path, head)

	for e := q.list.Front(); e != nil; e = e.Next() {
		if e != q.list.Front() {
			n, err := q.dequeue()
			if err == false {
				log.Fatalf("coundn't dequeue!!!\n")
			}
			path = append(path, n)
			fmt.Printf("path %+v\n", path)
		}
		node := e.Value.(*DS.GraphNode[T])
		if node.Value == key {
			return node, path, nil
		} else {
			for _, c := range node.Childs {
				q.enqueue(c)
			}
		}
	}
	return nil, nil, fmt.Errorf("coudn't find %+v\n", key)
}
