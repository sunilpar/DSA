package Test

import (
	"DSA/DataStructures"
	"testing"
)

func TestLinkedList(t *testing.T) {
	ll := &DS.LinkedList[int]{}

	// Test IsEmpty on an empty list
	if !ll.IsEmpty() {
		t.Errorf("Expected list to be empty")
	}

	// Test InsertAtFront
	ll.InsertAtFront(10)
	ll.InsertAtFront(20)
	if ll.Size() != 2 {
		t.Errorf("Expected size 2, got %d", ll.Size())
	}

	// Test Insert
	ll.Insert(30)
	if ll.Size() != 3 {
		t.Errorf("Expected size 3, got %d", ll.Size())
	}

	// Test Peek
	val, err := ll.Peek()
	if err != nil || val != 20 {
		t.Errorf("Expected head 20, got %d", val)
	}

	// Test PeekAtEnd
	val, err = ll.PeekAtEnd()
	if err != nil || val != 30 {
		t.Errorf("Expected tail 30, got %d", val)
	}

	// Test Delete
	err = ll.Delete()
	if err != nil || ll.Size() != 2 {
		t.Errorf("Expected size 2 after deleting head, got %d", ll.Size())
	}

	// Test DeleteAtEnd
	err = ll.DeleteAtEnd()
	if err != nil || ll.Size() != 1 {
		t.Errorf("Expected size 1 after deleting tail, got %d", ll.Size())
	}

	// Test deleting until empty
	err = ll.Delete()
	err = ll.Delete()
	if !ll.IsEmpty() {
		t.Errorf("Expected list to be empty after deleting all elements")
	}

	// Test deleting from empty list
	err = ll.Delete()
	if err == nil {
		t.Errorf("Expected error when deleting from empty list")
	}
}
