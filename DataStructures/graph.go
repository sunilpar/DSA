package DS

import (
	"fmt"
)

//should do this with adj matrix or adj list

type GraphNode[T any] struct {
	Value  T
	Childs []*GraphNode[T]
}

type Graph[T comparable] struct {
	Nodes *LinkedList[*GraphNode[T]]
}

func CreateGraph[T comparable]() *Graph[T] {
	return &Graph[T]{
		Nodes: &LinkedList[*GraphNode[T]]{},
	}
}

func (g *Graph[T]) Root(value T) error {
	if g.Nodes.Size() != 0 {
		return fmt.Errorf("can't add root on non empty graph size of graph:%v ", g.Nodes.Size())
	}
	node := &GraphNode[T]{Value: value}
	g.Nodes.InsertAtFront(node)
	return nil
}

func (g *Graph[T]) Insertafter(root T, value T) error {
	newNode := &GraphNode[T]{Value: value}
	gn := &GraphNode[T]{}
	for e := g.Nodes.Head; e != nil; e = e.Next {
		if e.Value.Value == root {
			gn = e.Value
			g.Nodes.Insert(newNode)
			g.Nodes.Tail.Prev = e
			gn.Childs = append(gn.Childs, newNode)
			return nil
		}
	}
	return fmt.Errorf("unknown parent node :%v\n", root)
}

func (g *Graph[T]) Delete(value T) error {
	if h := g.Nodes.Head.Value; h.Value == value || g.Nodes.Size() == 1 {
		err := g.DeleteRoot(value)
		if err != nil {
			return fmt.Errorf("%v", err)
		}
		return nil
	}
	for e := g.Nodes.Head; e.Next != nil; e = e.Next {
		if e.Next.Value.Value == value {
			for _, c := range e.Next.Value.Childs {
				for ea := g.Nodes.Head; ea != nil; ea = ea.Next {
					if ea.Value.Value == c.Value {
						ea.Prev = nil
					}
				}
			}
			e.Next = e.Next.Next
			return nil
		}
	}
	return fmt.Errorf("reached nil couldn't find :%v\n", value)
}

func (g *Graph[T]) DeleteRoot(value T) error {
	if h := g.Nodes.Head.Value; h.Value == value {
		for _, c := range h.Childs {
			for ea := g.Nodes.Head; ea != nil; ea = ea.Next {
				if ea.Value.Value == c.Value {
					ea.Prev = nil
				}
			}
		}
		g.Nodes.Head = g.Nodes.Head.Next
		return nil
	}
	return fmt.Errorf("couldn't match value of head to root:%v", value)
}

func (g *Graph[T]) Display() {
	fmt.Print("\nnodes graphs:vvv\n")
	for e := g.Nodes.Head; e != nil; e = e.Next {
		fmt.Printf("%v", e.Value.Value)
		fmt.Printf("->%+v ", e)
		fmt.Printf("%p", e)
		fmt.Printf("\n")
	}
}
