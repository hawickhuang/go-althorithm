package offer

import "testing"

func TestCQueue_AppendTail(t *testing.T) {
	cq := NewCQueue()
	cq.AppendTail(1)
	r := cq.DeleteHead()
	if r != 1 {
		t.Errorf("test DeleteHead error, want: %d, got %d", 1, r)
	}
	cq.AppendTail(2)
	cq.AppendTail(3)
	r = cq.DeleteHead()
	if r != 2 {
		t.Errorf("test DeleteHead error, want: %d, got %d", 2, r)
	}
	r = cq.DeleteHead()
	if r != 3 {
		t.Errorf("test DeleteHead error, want: %d, got %d", 3, r)
	}
	r = cq.DeleteHead()
	if r != -1 {
		t.Errorf("test DeleteHead error, want: %d, got %d", -1, r)
	}
}

