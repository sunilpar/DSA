package main

import (
	DS "DSA/DataStructures"
	"fmt"
	"log"
)

func main() {
	s := DS.Stack[int]{
		L: DS.LinkedList[int]{},
	}
	s.Push(10)
	s.Push(20)
	s.Push(30)

	v, err := s.Pop()
	if err != nil {
		log.Printf("error: %s\n", err)
	}
	fmt.Printf("poped :%v\n", v)

	v, err = s.Pop()
	if err != nil {
		log.Printf("error: %s\n", err)
	}
	fmt.Printf("poped :%v\n", v)

	v, err = s.Peek()
	if err != nil {
		log.Printf("error: %s\n", err)
	}
	fmt.Printf("peek value :%v\n", v)

	s.Display()
}
