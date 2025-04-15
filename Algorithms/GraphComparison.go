package Algo

//uses bfs to compare graph with any number of node
//and tells wether it's structurally and equal in value

import (
	DS "DSA/DataStructures"
	"container/list"
	"fmt"
	"log"
)

type queue1[T comparable] struct {
	list *list.List
}

type queue2[T comparable] struct {
	list *list.List
}

func newQueue1[T comparable]() *queue[T] {
	return &queue[T]{list: list.New()}
}
func newQueue2[T comparable]() *queue[T] {
	return &queue[T]{list: list.New()}
}

func (q *queue[T]) enqueue1(v *DS.GraphNode[T]) {
	q.list.PushBack(v)
}
func (q *queue[T]) enqueue2(v *DS.GraphNode[T]) {
	q.list.PushBack(v)
}

func (q *queue[T]) dequeue1() (*DS.GraphNode[T], bool) {
	if q.list.Len() == 0 {
		return nil, false
	}
	elem := q.list.Front()
	q.list.Remove(elem)
	return elem.Value.(*DS.GraphNode[T]), true
}
func (q *queue[T]) dequeue2() (*DS.GraphNode[T], bool) {
	if q.list.Len() == 0 {
		return nil, false
	}
	elem := q.list.Front()
	q.list.Remove(elem)
	return elem.Value.(*DS.GraphNode[T]), true
}

func CompareGraphs[T comparable](a *DS.Graph[T], b *DS.Graph[T]) (bool, error) {
	if a.Nodes.Size() == 0 || b.Nodes.Size() == 0 {
		return false, fmt.Errorf("graphs are empty!! \n")
	}
	head1 := a.Nodes.Head.Value
	head2 := b.Nodes.Head.Value
	q1 := newQueue1[T]()
	q2 := newQueue2[T]()
	q1.enqueue1(head1)
	q2.enqueue2(head2)

	e1 := q1.list.Front()
	e2 := q2.list.Front()
	for {
		if e1 == nil && e2 == nil {
			break
		} else if e1 == nil || e2 == nil {
			return false, nil
		}

		if e1 != q1.list.Front() && e2 != q2.list.Front() {
			n1, err := q1.dequeue1()
			if err == false {
				log.Fatalf("coundn't dequeue!!!:%v\n", err)
			}
			n2, err := q2.dequeue2()
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
