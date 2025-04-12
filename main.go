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
	g.Delete(2)
	g.Display()

}
