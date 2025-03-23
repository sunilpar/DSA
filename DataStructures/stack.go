package DS

import ()

type Stack[T any] struct {
	L LinkedList[T]
}

func (s *Stack[T]) Push(value T) {
	s.L.Insert(value)
}

func (s *Stack[T]) Pop() (T, error) {
	var zero T
	v, err := s.L.PeekAtEnd()
	if err != nil {
		return zero, err
	}
	err = s.L.DeleteAtEnd()
	if err != nil {
		return zero, err
	}
	return v, nil
}

func (s *Stack[T]) Peek() (T, error) {
	var zero T
	v, err := s.L.PeekAtEnd()
	if err != nil {
		return zero, err
	}
	return v, nil
}

func (s *Stack[T]) Display() {
	s.L.Display()
}
