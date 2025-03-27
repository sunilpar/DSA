package DS

import "fmt"

type Ringbuffer[T any] struct {
	Arr  []T
	Head int
	Tail int
	Size int
}

func (r *Ringbuffer[T]) EnqueueInFront(value T) {
	if ((r.Head-1)+r.Size)%r.Size == r.Tail {
		r.reallocate()
	}
	r.Head = ((r.Head - 1) + r.Size) % r.Size
	r.Arr[r.Head] = value
}

func (r *Ringbuffer[T]) Enqueue(value T) {
	if (r.Tail+1)%r.Size == r.Head {
		r.reallocate()
	}
	r.Arr[r.Tail] = value
	r.Tail = (r.Tail + 1) % r.Size
}

func (r *Ringbuffer[T]) Dequeue() error {
	if r.Tail == r.Head {
		return fmt.Errorf("cant dequeue empty array\n")
	}
	r.Tail = ((r.Tail - 1) + r.Size) % r.Size

	if r.Tail == r.Head {
		r.Head = r.Size / 2
		r.Tail = r.Size / 2
	}
	return nil
}

func (r *Ringbuffer[T]) DequeueFromFront() error {
	if r.Head == r.Tail {
		return fmt.Errorf("cant dequeue from empty array!!\n")
	}
	r.Head = (r.Head + 1) % r.Size
	if r.Head == r.Tail {
		r.Head = r.Size / 2
		r.Tail = r.Size / 2
	}
	return nil
}

func (r Ringbuffer[T]) Display() {
	if r.Tail != r.Head {
		i := r.Head
		fmt.Print("\n")
		for {
			fmt.Printf("%v:%v ", i, r.Arr[i])

			if (i+1)%r.Size == r.Tail {
				break
			}
			i = (i + 1) % r.Size
		}
	} else {
		fmt.Printf("Array is empty !!\n")
	}
}

func (r *Ringbuffer[T]) Len() int {
	if r.Tail == r.Head {
		return 0
	}
	t_h := (r.Tail - r.Head)
	if t_h < 0 {
		h_t := (r.Head - r.Tail)
		return (r.Size - h_t)
	}
	return (r.Tail - r.Head)
}

func (r *Ringbuffer[T]) reallocate() {
	nsize := r.Size * r.Size
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
