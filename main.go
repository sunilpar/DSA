package main

import (
	DS "DSA/DataStructures"
	"log"
)

func main() {
	g := DS.CreateAdjList[int]()
	g.AddNode(20)
	g.AddNode(30)
	err := g.AddEdge(20, 30, 0)
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}
	g.AddNode(10)
}
