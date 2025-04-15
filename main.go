package main

import (
	Algo "DSA/Algorithms"
	DS "DSA/DataStructures"
	"fmt"
	"log"
)

func main() {
	g := DS.CreateGraph[int]()
	g.Root(1)
	g.Insertafter(1, 2)
	g.Insertafter(1, 3)
	g.Insertafter(2, 4)
	g.Insertafter(2, 5)
	g.Insertafter(3, 6)
	g.Insertafter(3, 7)

	err := g.Delete(2)
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}

	g.Display()

	h := DS.CreateGraph[int]()
	h.Root(1)
	h.Insertafter(1, 2)
	h.Insertafter(1, 3)
	h.Insertafter(2, 4)
	h.Insertafter(2, 5)
	h.Insertafter(3, 6)
	h.Insertafter(3, 8)

	err = h.Delete(2)
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}

	h.Display()

	s, err := Algo.CompareGraphs(g, h)
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}
	fmt.Printf("\nis the graph same:%v\n", s)

}
