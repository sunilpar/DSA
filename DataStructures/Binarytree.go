package DS

import "fmt"

type Ordered interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64 | ~string
}

type BinaryNode[T Ordered] struct {
	Value T
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

// need to delete and replace with the largest on left or smallest in right
// case 1 no child
// case 2 one child
// case 3 2 child and need to find largest element on right side to replace on parent node
func (g *Btree[T]) Delete() {
}

func (g *Btree[T]) Insert(n *BinaryNode[T], v T) error {
	if n.Left == nil && n.Right == nil {
		Newnode := &BinaryNode[T]{Value: v}
		if v <= n.Value {
			n.Left = Newnode
			fmt.Printf("%v:<--\n", v)
		} else {
			n.Right = Newnode
			fmt.Printf("-->:%v\n", v)
		}
		g.Nodes.Insert(Newnode)
		return nil
	} else {
		if v <= n.Value {
			if n.Left == nil {
				Newnode := &BinaryNode[T]{Value: v}
				n.Left = Newnode
				fmt.Printf("%v:<--\n", v)
				g.Nodes.Insert(Newnode)
				return nil
			}
			return g.Insert(n.Left, v)
		} else {
			if n.Right == nil {
				Newnode := &BinaryNode[T]{Value: v}
				n.Right = Newnode
				fmt.Printf("-->:%v\n", v)
				g.Nodes.Insert(Newnode)
				return nil
			}
			return g.Insert(n.Right, v)
		}
	}
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
