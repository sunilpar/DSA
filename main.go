package main

import (
	DS "DSA/DataStructures"
	"fmt"
	"log"
	"time"
)

func main() {
	g := DS.CreateGraph[int]()
	s := time.Now()
	g.Root(1)
	e := time.Since(s)
	fmt.Printf("\ninsettion time taken:%v\n", e)

	s = time.Now()
	g.Insertafter(1, 2)
	e = time.Since(s)
	fmt.Printf("\ninsettion time taken:%v\n", e)

	s = time.Now()
	g.Insertafter(2, 4)
	e = time.Since(s)
	fmt.Printf("\ninsettion time taken:%v\n", e)

	s = time.Now()
	g.Insertafter(2, 5)
	e = time.Since(s)
	fmt.Printf("\ninsettion time taken:%v\n", e)

	s = time.Now()
	g.Insertafter(3, 6)
	e = time.Since(s)
	fmt.Printf("\ninsettion time taken:%v\n", e)

	s = time.Now()
	g.Insertafter(3, 7)
	e = time.Since(s)
	fmt.Printf("\ninsettion time taken:%v\n", e)

	s = time.Now()
	err := g.Delete(2)
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}
	e = time.Since(s)
	fmt.Printf("\ndeletion time taken:%v\n", e)

	g.Display()

}
