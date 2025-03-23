package main

import (
	DS "DSA/DataStructures"
	"fmt"
	"log"
)

func main() {
	q := DS.Queue[int]{
		L: DS.LinkedList[int]{},
	}

	v, err := q.Enqueue(10)
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	} else {
		fmt.Printf("val inserted :%d \n", v)
	}

	v, err = q.Enqueue(20)
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	} else {
		fmt.Printf("val inserted :%d \n", v)
	}

	v, err = q.Enqueue(30)
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	} else {
		fmt.Printf("val inserted :%d \n", v)
	}

	v, err = q.Enqueue(40)
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	} else {
		fmt.Printf("val inserted :%d \n", v)
	}

	err = q.Dequeue()
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}

	q.Display()

}
