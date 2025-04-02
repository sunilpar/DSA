package DS

import (
	DS "DSA/DataStructures"
	"testing"
)

func TestRingBuffer(t *testing.T) {
	rb := DS.Ringbuffer[int]{Arr: make([]int, 5), Size: 5, Head: 2, Tail: 2}

	// Test Initial State
	if rb.Len() != 0 {
		t.Errorf("Expected length 0, got %d", rb.Len())
	}

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

	// Test Wrap Around
	rb.Enqueue(40)
	rb.Enqueue(50)
	rb.Enqueue(60) // This should trigger reallocation

	if rb.Len() != 5 {
		t.Errorf("Expected length 5 after wrap-around insertions, got %d", rb.Len())
	}

	// Test Reallocation
	for i := 0; i < 10; i++ {
		rb.Enqueue(i * 10)
	}

	if rb.Size < 15 { // Size should have increased significantly
		t.Errorf("Expected size > 15 after multiple reallocations, got %d", rb.Size)
	}

	// Test Dequeue Until Empty
	for i := 0; i < rb.Len(); i++ {
		rb.Dequeue()
	}

	if rb.Len() != 0 {
		t.Errorf("Expected length 0 after dequeuing all elements, got %d", rb.Len())
	}

	// Test Dequeue from Empty Buffer
	err = rb.Dequeue()
	if err == nil {
		t.Errorf("Expected error on dequeue from empty buffer, got nil")
	}

	// Test DequeueFromFront on Empty Buffer
	err = rb.DequeueFromFront()
	if err == nil {
		t.Errorf("Expected error on front dequeue from empty buffer, got nil")
	}
}
