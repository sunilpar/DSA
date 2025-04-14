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
	f, p, err := Algo.DFS(g, 2)
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}
	fmt.Printf("path:")
	for _, v := range p {
		fmt.Printf("%v ", v.Value)
	}
	fmt.Print("\n")
	fmt.Printf("f.value:%v\n", f.Value)

	g.Display()

}
