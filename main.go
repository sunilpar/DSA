package main

import (
	DS "DSA/DataStructures"
	"log"
)

func main() {
	g := DS.CreateAdjList[int]()
	err := g.AddNode(20)
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}
	err = g.AddNode(30)
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}
	err = g.AddEdge(20, 30, 0)
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}
	err = g.DeleteNode(20)
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}
}
