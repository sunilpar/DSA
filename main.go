package main

import (
	DS "DSA/DataStructures"
	"fmt"
	"log"
)

func main() {

	l := DS.LinkedList[int]{}
	l.InsertAtFront(10)
	l.Insert(20)
	l.Insert(30)

	val, err := l.Peek()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("peek :%d\n", val)
	}

	endval, err := l.PeekAtEnd()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("peek :%d\n", endval)
	}

	fmt.Printf("isEmpty:%v\n", l.IsEmpty())

	fmt.Printf("size :%d\n", l.Size())

	l.Delete()
	l.Delete()
	l.Delete()
	fmt.Printf("isEmpty:%v\n", l.IsEmpty())
	l.Delete()

	l.Display()

}
