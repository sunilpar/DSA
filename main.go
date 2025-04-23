package main

import (
	DS "DSA/DataStructures"
	"fmt"
	"log"
)

func main() {
	g := DS.CreateBtree[string]()
	H, err := g.Root("sunil")
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}
	fmt.Printf("<%v>\n", H.Value)

	g.Insert(H, "left hand")
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}
	g.Insert(H, "righthand")
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}
	g.Insert(H, "body")
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}
	g.Insert(H, "left leg")
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}
	g.Insert(H, "right leg")
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}

	g.Display()

}
