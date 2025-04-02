package DS

import (
	DS "DSA/DataStructures"
	"testing"
)

func TestRingBuffer(t *testing.T) {
	rb := DS.Ringbuffer[int]{Arr: make([]int, 5), Size: 5, Head: 2, Tail: 2}

	// Test Enqueue
	rb.Enqueue(10)
	rb.Enqueue(20)
	rb.Enqueue(30)

	if rb.Len() != 3 {
		t.Errorf("Expected length 3, got %d", rb.Len())
	}

	// Test EnqueueInFront
	rb.EnqueueInFront(5)
	if rb.Arr[rb.Head] != 5 {
		t.Errorf("Expected front element 5, got %d", rb.Arr[rb.Head])
	}

	// Test Dequeue
	err := rb.Dequeue()
	if err != nil {
		t.Errorf("Unexpected error on dequeue: %v", err)
	}

	if rb.Len() != 3 {
		t.Errorf("Expected length 3 after dequeue, got %d", rb.Len())
	}

	// Test DequeueFromFront
	err = rb.DequeueFromFront()
	if err != nil {
		t.Errorf("Unexpected error on front dequeue: %v", err)
	}

	if rb.Len() != 2 {
		t.Errorf("Expected length 2 after front dequeue, got %d", rb.Len())
	}
}
