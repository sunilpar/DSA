package DS

import "fmt"

type Queue struct {
	L LinkedList
}

func (Q Queue) DisplayQ() {
	fmt.Printf("queue > %s", Q.L.DisplayL())
}
