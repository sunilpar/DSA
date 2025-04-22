package main

import (
	DS "DSA/DataStructures"
	"fmt"
	"log"
)

func main() {
	g := DS.CreateBtree[int]()
	H, err := g.Root(20)
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}
	fmt.Printf("<%v>\n", H.Value)

	g.Insert(H, 10)
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}
	g.Insert(H, 30)
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}

	g.Insert(H, 0)
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}
	g.Insert(H, 19)
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}
	g.Insert(H, 11)
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}

	g.Insert(H, 25)
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}
	g.Insert(H, 40)
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}

	err = g.Delete(H, 0)
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}

	g.Insert(H, 0)
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}
	g.Display()

}
