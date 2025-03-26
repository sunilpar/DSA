package main

import (
	DS "DSA/DataStructures"
)

func main() {
	r := DS.Ringbuffer[int]{
		Arr:  make([]int, 4),
		Head: 4 / 2,
		Tail: 4 / 2,
		Size: 4,
	}

	r.EnqueueInFront(10)
	r.Enqueue(20)
	r.Enqueue(30)
	r.Enqueue(40)
	r.Enqueue(50)
	r.Enqueue(60)
	r.Enqueue(70)
	r.Enqueue(80)
	r.Enqueue(90)
	r.Enqueue(100)

	r.Display()
}
