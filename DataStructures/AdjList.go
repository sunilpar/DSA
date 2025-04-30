package DS

type ListNode[T any] struct {
	Value  T
	Childs []*ListEdge[T]
}

type ListEdge[T any] struct {
	To     T
	Weight int
}

func CreateAdjList[T comparable]() *ListNode[T] {
	return &ListNode[T]{}
}
