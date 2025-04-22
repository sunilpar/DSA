package DS

import (
	"fmt"
)

type Node[T any] struct {
	Value T
	Next  *Node[T]
	Prev  *Node[T]
}

type LinkedList[T any] struct {
	Head *Node[T]
	Tail *Node[T]
}

func (l *LinkedList[T]) InsertAtFront(Value T) *Node[T] {
	if l.Head == nil {
		newNode := &Node[T]{Value: Value, Next: l.Head, Prev: nil}
		l.Head = newNode
		l.Tail = newNode
		return newNode

	} else {

		newNode := &Node[T]{Value: Value, Next: l.Head, Prev: nil}
		l.Head.Prev = newNode
		l.Head = newNode
		return newNode
	}
}

func (l *LinkedList[T]) Insert(Value T) *Node[T] {
	if l.Head == nil {
		newNode := &Node[T]{Value: Value, Next: l.Head, Prev: nil}
		l.Tail = newNode
		l.Head = newNode
		return newNode

	} else {
		newNode := &Node[T]{Value: Value, Next: nil, Prev: l.Tail}
		l.Tail.Next = newNode
		l.Tail = newNode
		return newNode
	}
}

func (l *LinkedList[T]) Delete() error {
	//del from front
	if l.Head == nil {
		return fmt.Errorf("cant delete on empty list!!!")
	} else {
		if l.Head.Next == nil {
			l.Head = nil
			l.Tail = nil
		} else {
			l.Head.Next.Prev = nil
			l.Head = l.Head.Next
		}
	}
	return nil
}

func (l *LinkedList[T]) DeleteAtEnd() error {
	//del from end
	if l.Tail == nil {
		return fmt.Errorf("cant delete on empty list!!!")
	} else {
		if l.Tail.Prev == nil {
			l.Tail = nil
			l.Head = nil
		} else {
			l.Tail.Prev.Next = nil
			l.Tail = l.Tail.Prev
		}
	}
	return nil
}

// func (l *LinkedList[T]) DeleteM(m T) error {
// 	for e := l.Head; e != nil; e = e.Next {
// 		if e.Value == m {
// 			e.Prev.Next = e.Next
// 			e.Next.Prev = e.Prev
// 			e = nil
// 			return nil
// 		}
// 	}
// 	return fmt.Errorf("value %v not found in the list", m)
// }

func (l LinkedList[T]) Size() int {
	current := l.Head
	count := 0

	for current != nil {
		count++
		current = current.Next
	}
	return count
}

func (l *LinkedList[T]) Peek() (T, error) {
	var zero T
	if l.Head == nil {
		return zero, fmt.Errorf("cant peek on empty list!!!\n")
	}
	return l.Head.Value, nil
}

func (l *LinkedList[T]) PeekAtEnd() (T, error) {
	var zero T
	if l.Tail == nil {
		return zero, fmt.Errorf("cant peek on empty list!!!\n")
	}
	return l.Tail.Value, nil
}

func (l *LinkedList[T]) IsEmpty() bool {
	if l.Tail == nil {
		return true
	}
	return false
}

func (l LinkedList[T]) Display() {
	current := l.Head
	for current != nil {
		fmt.Printf("%v ", current.Value)
		current = current.Next
	}
}
