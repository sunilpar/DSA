package DS

import "log"

type Queue[T any] struct {
	L LinkedList[T]
}

func (q *Queue[T]) Enqueue(value T) (T, error) {
	var v, zero T
	q.L.Insert(value)
	v, err := q.L.Peek()
	if err != nil {
		log.Fatalf("error: %s\n", err)
		return zero, err
	}
	return v, nil
}

func (q *Queue[T]) Dequeue() error {
	err := q.L.Delete()
	if err != nil {
		log.Fatalf("error: %s\n", err)
		return err
	}
	return nil
}

func (q Queue[T]) Display() {
	q.L.Display()
}
