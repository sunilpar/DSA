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
	g.Display()
	gn, path, err := Algo.BFS(g, 1)
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}
	g.Insertafter(gn, 2)
	fmt.Printf("\ngn.value :%v\ngn.chldes.values:", gn.Value)
	for _, v := range gn.Childs {
		fmt.Printf("%v ", v.Value)
	}
	fmt.Printf("\n")

	fmt.Printf("\npath.values:")
	for _, v := range path {
		fmt.Printf("%v ", v.Value)
	}
	fmt.Printf("\n")

	gn, path, err = Algo.BFS(g, 1)
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}
	g.Insertafter(gn, 3)
	fmt.Printf("\ngn.value :%v\ngn.chldes.values:", gn.Value)
	for _, v := range gn.Childs {
		fmt.Printf("%v ", v.Value)
	}
	fmt.Printf("\n")

	fmt.Printf("\npath.values:")
	for _, v := range path {
		fmt.Printf("%v ", v.Value)
	}
	fmt.Printf("\n")
}
