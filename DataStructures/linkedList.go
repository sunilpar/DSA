package DS

import (
	"fmt"
)

type Node[T any] struct {
	value T
	next  *Node[T]
	prev  *Node[T]
}

type LinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
}

func (l *LinkedList[T]) InsertAtFront(value T) {
	if l.head == nil {
		newNode := &Node[T]{value: value, next: l.head, prev: nil}
		l.head = newNode
		l.tail = newNode
		fmt.Printf("inserted: %v and node is : %p\n", newNode.value, l.head)

	} else {

		newNode := &Node[T]{value: value, next: l.head, prev: nil}
		l.head.prev = newNode
		l.head = newNode
		fmt.Printf("inserted: %v and node is : %p\n", newNode.value, l.head)
	}
}

func (l *LinkedList[T]) Insert(value T) {
	if l.head == nil {
		newNode := &Node[T]{value: value, next: l.head, prev: nil}
		l.tail = newNode
		l.head = newNode
		fmt.Printf("inserted: %v and node is : %p\n", newNode.value, l.tail)

	} else {
		newNode := &Node[T]{value: value, next: nil, prev: l.tail}
		l.tail.next = newNode
		l.tail = newNode
		fmt.Printf("inserted: %v and node is : %p\n", newNode.value, l.tail)
	}
}

// del functions

func (l *LinkedList[T]) Delete() error {
	//del from front
	if l.head == nil {
		return fmt.Errorf("cant delete on empty list!!!")
	} else {
		if l.head.next == nil {
			l.head = nil
			l.tail = nil
		} else {
			l.head.next.prev = nil
			l.head = l.head.next
		}

		//		fmt.Printf("new head is :%p\n", l.head)
	}
	return nil
}

func (l *LinkedList[T]) DeleteAtEnd() error {
	//del from end
	if l.tail == nil {
		return fmt.Errorf("cant delete on empty list!!!")
	} else {
		if l.tail.prev == nil {
			l.tail = nil
			l.head = nil
		} else {
			l.tail.prev.next = nil
			l.tail = l.tail.prev
		}

		//	fmt.Printf("new tail is :%p\n", l.tail)
	}
	return nil
}

func (l LinkedList[T]) Size() int {
	current := l.head
	count := 0

	for current != nil {
		count++
		current = current.next
	}
	return count
}

func (l *LinkedList[T]) Peek() (T, error) {
	var zero T
	if l.head == nil {
		return zero, fmt.Errorf("cant peek on empty list!!!\n")
	}
	return l.head.value, nil
}

func (l *LinkedList[T]) PeekAtEnd() (T, error) {
	var zero T
	if l.tail == nil {
		return zero, fmt.Errorf("cant peek on empty list!!!\n")
	}
	return l.tail.value, nil
}

func (l *LinkedList[T]) IsEmpty() bool {
	if l.tail == nil {
		return true
	}
	return false
}

func (l LinkedList[T]) Display() {
	current := l.head
	fmt.Printf("head was : %v \n", current)

	for current != nil {
		fmt.Printf("%v ", current.value)
		current = current.next
	}

}
