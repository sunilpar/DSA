package DS

import "fmt"

type Ordered interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64 | ~string
}

type BinaryNode[T Ordered] struct {
	Value T
	Prev  *BinaryNode[T]
	Right *BinaryNode[T]
	Left  *BinaryNode[T]
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
		return nil, fmt.Errorf("cannot search into a nil node")
	}
	if v == n.Value {
		return n, nil
	}
	if v <= n.Value {
		fmt.Printf("going left <--%v\n", n.Value)
		return g.Search(n.Left, v)
	} else {
		fmt.Printf("going right -->%v\n", n.Value)
		return g.Search(n.Right, v)
	}
}

// need to delete and replace with the largest on left or smallest in right
// case 1 no child
// case 2 one child
// case 3 2 child and need to find largest element on right side to replace on parent node
func (g *Btree[T]) Delete() {
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
