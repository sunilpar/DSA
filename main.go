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
	g.InsertAFter(gn, 2)
	fmt.Printf("path was :%v\n", path)
	fmt.Printf("gn was :%+v\n", gn)

}
