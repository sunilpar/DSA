package DS

import "fmt"

//binary tree with left and right as children
//binarytreenode as node type
//need to write search insertion deletion display with dfs

type BinaryNode[T any] struct {
	Value T
	Right *BinaryNode[T]
	Left  *BinaryNode[T]
}

type Btree[T comparable] struct {
	Nodes *LinkedList[*BinaryNode[T]]
}

func CreateBtree[T comparable]() *Btree[T] {
	return &Btree[T]{
		Nodes: &LinkedList[*BinaryNode[T]]{},
	}
}

func (g *Btree[T]) Root(value T) error {
	if g.Nodes.Size() != 0 {
		return fmt.Errorf("can't add root on non empty graph size of graph:%v ", g.Nodes.Size())
	}
	node := &BinaryNode[T]{Value: value}
	g.Nodes.InsertAtFront(node)
	return nil
}

func (g *Btree[T]) InsertRight(root T, value T) error {
	newNode := &BinaryNode[T]{Value: value}
	gn := &BinaryNode[T]{}
	for e := g.Nodes.Head; e != nil; e = e.Next {
		if e.Value.Value == root {
			gn = e.Value
			g.Nodes.Insert(newNode)
			g.Nodes.Tail.Prev = e
			gn.Right = newNode
			return nil
		}
	}
	return fmt.Errorf("unknown parent node :%v\n", root)

}

func (g *Btree[T]) InsertLeft(root T, value T) error {
	newNode := &BinaryNode[T]{Value: value}
	gn := &BinaryNode[T]{}
	for e := g.Nodes.Head; e != nil; e = e.Next {
		if e.Value.Value == root {
			gn = e.Value
			g.Nodes.Insert(newNode)
			g.Nodes.Tail.Prev = e
			gn.Left = newNode
			return nil
		}
	}
	return fmt.Errorf("unknown parent node :%v\n", root)

}
func (g *Btree[T]) Display() {
	fmt.Print("\nnodes graphs:vvv\n")
	for e := g.Nodes.Head; e != nil; e = e.Next {
		fmt.Printf("%v", e.Value.Value)
		fmt.Printf("->%+v ", e)
		fmt.Printf("%p", e)
		fmt.Printf("\n")
	}
}
