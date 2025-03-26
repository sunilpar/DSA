package main

import (
	DS "DSA/DataStructures"
)

func main() {
	r := DS.Ringbuffer[int]{
		Arr:  make([]int, 4), // Initialize the array
		Head: 4 / 2,
		Tail: 4 / 2,
		Size: 4,
	}

	r.Enqueue(10)
	r.Enqueue(20)
	r.Enqueue(30)
	r.Enqueue(40)
	r.Enqueue(50)
	r.Enqueue(60)

	r.Display()
}
