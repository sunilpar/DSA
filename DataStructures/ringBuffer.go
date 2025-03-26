package DS

import "fmt"

type Ringbuffer[T any] struct {
	Arr  []T
	Head int
	Tail int
	Size int
}

func (r *Ringbuffer[T]) EnqueueInFront(value T) {
	prevh := r.Head
	prevt := r.Tail
	if ((r.Head-1)+r.Size)%r.Size == r.Tail {
		r.reallocate()
	}
	r.Head = ((r.Head - 1) + r.Size) % r.Size
	r.Arr[r.Head] = value
	fmt.Printf("\n head was:%v\n tail was:%v \n inserted %v at :[%v:%v]\n", prevh, prevt, value, r.Head, r.Tail)
}

func (r *Ringbuffer[T]) Enqueue(value T) {
	prevh := r.Head
	prevt := r.Tail
	if (r.Tail+1)%r.Size == r.Head {
		r.reallocate()
	}
	r.Arr[r.Tail] = value
	r.Tail = (r.Tail + 1) % r.Size
	fmt.Printf("\n head was:%v\n tail was:%v \n inserted %v at :[%v:%v]\n", prevh, prevt, value, r.Head, r.Tail)
}

func (r Ringbuffer[T]) Display() {
	fmt.Printf("\n display[%v:%v]\n", r.Head, r.Tail)

	i := r.Head
	for {
		fmt.Printf("%v:%v ", i, r.Arr[i])

		if (i+1)%r.Size == r.Tail {
			break
		}
		i = (i + 1) % r.Size
	}
}

func (r *Ringbuffer[T]) reallocate() {
	fmt.Printf("\nbefore allocation %v[%v:%v]\n", r.Size, r.Head, r.Tail)
	nsize := r.Size * 2
	newa := make([]T, nsize)

	j := nsize / 2
	i := r.Head
	for {
		newa[j] = r.Arr[i]
		if i == r.Tail {
			break
		}
		j++
		i = (i + 1) % r.Size
	}
	r.Head = nsize / 2
	r.Tail = j
	r.Arr = newa
	r.Size = nsize
	fmt.Printf("after allocation %v[%v:%v]\n", nsize, r.Head, r.Tail)
}
