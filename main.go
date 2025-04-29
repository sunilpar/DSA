package main

import (
	DS "DSA/DataStructures"
	"log"
)

func main() {
	g := DS.CreateGraph[string]()
	g.Root("sunil")
	g.Insertafter("sunil", "rhand")
	g.Insertafter("sunil", "lhand")
	g.Insertafter("sunil", "body")

	err := g.Delete("body")
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}

	g.Display()

}
