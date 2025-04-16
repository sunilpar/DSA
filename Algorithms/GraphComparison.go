package Algo

import (
	DS "DSA/DataStructures"
	"fmt"
	"log"
)

func CompareGraphs[T comparable](a *DS.Graph[T], b *DS.Graph[T]) (bool, error) {
	if a.Nodes.Size() == 0 || b.Nodes.Size() == 0 {
		return false, fmt.Errorf("graphs are empty!! \n")
	}
	head1 := a.Nodes.Head.Value
	head2 := b.Nodes.Head.Value
	q1 := newQueue[T]()
	q2 := newQueue[T]()
	q1.enqueue(head1)
	q2.enqueue(head2)

	e1 := q1.list.Front()
	e2 := q2.list.Front()
	for {
		if e1 != q1.list.Front() && e2 != q2.list.Front() {
			n1, err := q1.dequeue()
			if err == false {
				log.Fatalf("coundn't dequeue!!!:%v\n", err)
			}
			n2, err := q2.dequeue()
			if err == false {
				log.Fatalf("coundn't dequeue!!!:%v\n", err)
			}
			if n1.Value != n2.Value {
				return false, nil
			}
		} else {
			if e1.Value.(*DS.GraphNode[T]).Value != e2.Value.(*DS.GraphNode[T]).Value {
				return false, nil
			}
		}
		if e1 == nil && e2 == nil {
			break
		} else if e1 == nil || e2 == nil {
			return false, nil
		}

		node1 := e1.Value.(*DS.GraphNode[T])
		node2 := e2.Value.(*DS.GraphNode[T])
		for _, c := range node1.Childs {
			q1.enqueue(c)
		}
		for _, c := range node2.Childs {
			q2.enqueue(c)
		}

		e1 = e1.Next()
		e2 = e2.Next()
	}
	return true, nil
}
