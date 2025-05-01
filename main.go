package main

import (
	DS "DSA/DataStructures"
)

func main() {
	g := DS.CreateAdjList[int]()
	g.AddNode(20)
	g.AddNode(30)
	g.AddNode(10)
}
