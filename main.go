package main

import (
	DS "DSA/DataStructures"
	"log"
)

func main() {
	g := DS.CreateBtree[int]()
	err := g.Root(1)
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}
	g.InsertLeft(1, 2)
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}
	err = g.InsertLeft(0, 2)
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}
	g.InsertLeft(2, 4)
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}
	g.InsertRight(2, 5)
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}
	g.InsertRight(1, 3)
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}
	g.InsertLeft(3, 6)
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}
	g.InsertRight(3, 7)
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}

	g.Display()

}
