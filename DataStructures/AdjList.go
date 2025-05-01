package DS

import (
	"fmt"
	"log"
)

type Adjlist[T comparable] struct {
	Nodes []ListNode[T]
}
type ListNode[T comparable] struct {
	Value  T
	Childs []ListEdge[T]
}

type ListEdge[T comparable] struct {
	To     int
	Weight int
}

func CreateAdjList[T comparable]() *Adjlist[T] {
	return &Adjlist[T]{}
}

func (g *Adjlist[T]) rootList(value T) error {
	if len(g.Nodes) != 0 {
		return fmt.Errorf("graph already has other nodes!!")
	}
	r := ListNode[T]{Value: value}
	g.Nodes = append(g.Nodes, r)
	return nil
}

func (g *Adjlist[T]) AddNode(value T) error {
	if len(g.Nodes) == 0 {
		err := g.rootList(value)
		if err != nil {
			log.Fatalf("error: %s\n", err.Error())
		}
	}
	n := ListNode[T]{Value: value}
	g.Nodes = append(g.Nodes, n)
	return nil
}

func (g *Adjlist[T]) AddEdge(p, c T, w int) error {
	pi := -1
	ci := -1
	for i, v := range g.Nodes {
		if v.Value == p {
			pi = i
		}
		if v.Value == c {
			ci = i
		}
		if ci != -1 && pi != -1 {
			break
		}
	}
	if ci == -1 || pi == -1 {
		return fmt.Errorf("coundn't find child and parent :%v %v", pi, ci)
	}
	C := ListEdge[T]{To: ci, Weight: w}
	g.Nodes[pi].Childs = append(g.Nodes[pi].Childs, C)
	fmt.Printf("p:%v \n", g.Nodes[pi])
	return nil
}
