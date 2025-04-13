package main

import (
	DS "DSA/DataStructures"
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
	//
	// err := g.Delete(1)
	// if err != nil {
	// 	log.Fatalf("error: %s\n", err.Error())
	// }
	//
	// g.DeleteRoot(2)
	// if err != nil {
	// 	log.Fatalf("error: %s\n", err.Error())
	// }
	//
	// g.DeleteRoot(3)
	// if err != nil {
	// 	log.Fatalf("error: %s\n", err.Error())
	// }
	// err = g.Delete(4)
	// if err != nil {
	// 	log.Fatalf("error: %s\n", err.Error())
	// }
	//
	// err = g.Delete(5)
	// if err != nil {
	// 	log.Fatalf("error: %s\n", err.Error())
	// }
	//
	// err = g.Delete(6)
	// if err != nil {
	// 	log.Fatalf("error: %s\n", err.Error())
	// }
	//
	// err = g.Delete(7)
	// if err != nil {
	// 	log.Fatalf("error: %s\n", err.Error())
	// }

	g.Display()

}
