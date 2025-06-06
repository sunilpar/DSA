package DS

import (
	"fmt"
	"log"
)

type Ordered interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64 | ~string
}

type BinaryNode[T Ordered] struct {
	Value T
	Prev  *BinaryNode[T]
	Left  *BinaryNode[T]
	Right *BinaryNode[T]
	// ListNode *Node[*BinaryNode[T]]
}

type Btree[T Ordered] struct {
	Nodes *LinkedList[*BinaryNode[T]]
}

func CreateBtree[T Ordered]() *Btree[T] {
	return &Btree[T]{
		Nodes: &LinkedList[*BinaryNode[T]]{},
	}
}

func (g *Btree[T]) Root(value T) (*BinaryNode[T], error) {
	if g.Nodes.Size() != 0 {
		return nil, fmt.Errorf("can't add root on non empty graph size of graph:%v ", g.Nodes.Size())
	}
	node := &BinaryNode[T]{Value: value}
	g.Nodes.InsertAtFront(node)
	return g.Nodes.Head.Value, nil
}

func (g *Btree[T]) Search(n *BinaryNode[T], v T) (*BinaryNode[T], error) {
	if n == nil {
		return nil, fmt.Errorf("couldn't find the node reached nil!!")
	}
	if v == n.Value {
		return n, nil
	}
	if v <= n.Value {
		return g.Search(n.Left, v)
	} else {
		return g.Search(n.Right, v)
	}
}

func (g *Btree[T]) Delete(n *BinaryNode[T], v T) error {
	dn, err := g.Search(n, v)
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}
	if dn.Left == nil || dn.Right == nil {
		if dn.Right == nil && dn.Left == nil {
			//0 child
			if dn.Prev.Right == dn {
				dn.Prev.Right = nil
			} else {
				dn.Prev.Left = nil
			}
			fmt.Printf("deleted node with no child:%v\n", dn.Value)
			return nil
		}
		//1 child
		c := &BinaryNode[T]{}
		if dn.Right == nil {
			c = dn.Left
		} else {
			c = dn.Right
		}
		if dn.Prev.Right == dn {
			dn.Prev.Right = c
			c.Prev = dn.Prev
		} else {
			dn.Prev.Left = c
			c.Prev = dn.Prev
		}
		fmt.Printf("deleted node with 1 child:%v\n", dn.Value)
		return nil
	}
	//2 child
	c, err := g.SearchLargeChild(dn.Left)
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}
	dn.Value = c.Value
	g.Delete(c, c.Value)
	return nil
}

func (g *Btree[T]) SearchLargeChild(dn *BinaryNode[T]) (*BinaryNode[T], error) {
	if dn == nil {
		return nil, fmt.Errorf("nil node")
	}
	if dn.Right == nil {
		return dn, nil
	}
	return g.SearchLargeChild(dn.Right)
}

func (g *Btree[T]) Insert(n *BinaryNode[T], v T) error {
	if n == nil {
		return fmt.Errorf("cannot insert into a nil node")
	}

	var child **BinaryNode[T]
	if v <= n.Value {
		child = &n.Left
	} else {
		child = &n.Right
	}

	if *child == nil {
		newNode := &BinaryNode[T]{
			Value: v,
			Prev:  n,
		}
		*child = newNode
		g.Nodes.Insert(newNode)
		return nil
	}

	return g.Insert(*child, v)
}

func (g *Btree[T]) Display() {
	fmt.Print("\nnodes graphs:vvv\n")
	for e := g.Nodes.Head; e != nil; e = e.Next {
		fmt.Printf("%v", e.Value.Value)
		fmt.Printf("->%+v \t", e.Value)
		fmt.Printf("%p", e.Value)
		fmt.Printf("\n")
	}
}
