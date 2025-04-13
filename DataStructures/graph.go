package DS

import (
	"container/list"
	"fmt"
	"log"
)

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
	gn, _, err := BFS(g, root)
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}
	newNode := &GraphNode[T]{Value: value}
	for e := g.Nodes.Head; e != nil; e = e.Next {
		if e.Value == gn {
			g.Nodes.Insert(newNode)
			g.Nodes.Tail.Prev = e
		}
	}
	gn.Childs = append(gn.Childs, newNode)
	return nil
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

type queue[T comparable] struct {
	list *list.List
}

func newQueue[T comparable]() *queue[T] {
	return &queue[T]{list: list.New()}
}

func (q *queue[T]) enqueue(v *GraphNode[T]) {
	q.list.PushBack(v)
}

func (q *queue[T]) dequeue() (*GraphNode[T], bool) {
	if q.list.Len() == 0 {
		return nil, false
	}
	elem := q.list.Front()
	q.list.Remove(elem)
	return elem.Value.(*GraphNode[T]), true
}

func BFS[T comparable](graph *Graph[T], key T) (*GraphNode[T], []*GraphNode[T], error) {
	if graph.Nodes.Size() == 0 {
		return nil, nil, fmt.Errorf("Graph is empty:%v\n", graph)
	}
	elem := graph.Nodes.Head
	head := elem.Value
	q := newQueue[T]()
	q.enqueue(head)
	path := []*GraphNode[T]{}
	path = append(path, head)
	for {
		node, ok := q.dequeue()
		if !ok {
			break
		}
		path = append(path, node)
		if node.Value == key {
			return node, path, nil
		}
		for _, c := range node.Childs {
			q.enqueue(c)
		}
	}
	return nil, nil, fmt.Errorf("coudn't find %+v\n", key)
}
