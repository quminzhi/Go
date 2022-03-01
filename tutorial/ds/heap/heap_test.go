package heap

import (
	"testing"
)

func TestInitHeap(t *testing.T) {
	vec := []int{1, 2, 3, 4, 5}
	var h Heap
	if err := h.InitHeap(vec); err != nil {
		t.Fatalf("fail to init a heap\n")
	}
}

func TestHeapProperty(t *testing.T) {
	vec := []int{5, 3, 7, 11, 2}
	var h Heap
	if err := h.InitHeap(vec); err != nil {
		t.Fatalf("fail to init a heap\n")
	}

	// test size method
	if size, err := h.Size(); err == nil {
		if size != 5 {
			t.Fatalf("wrong size, expected size to be 5")
		}
	} else {
		t.Fatalf("failed to invoke h.Size()")
	}

	// test top method
	if top, err := h.Top(); err == nil {
		if top != 2 {
			t.Fatalf("wrong top, expected top to be 2")
		}
	} else {
		t.Fatalf("failed to invoke h.Top()")
	}
}

func TestPushOpt(t *testing.T) {
	vec := []int{5, 3, 7, 11, 2}
	var h Heap
	if err := h.InitHeap(vec); err != nil {
		t.Fatalf("failed to init a heap\n")
	}
	h.Push(1)
	if top, err := h.Top(); err != nil {
		t.Fatalf("failed to invoke Top()")
	} else {
		if top != 1 {
			t.Fatalf("got %v, but expected to be 1 after pushing 1", top)
		}
	}
	h.Push(2)
	if top, err := h.Top(); err != nil {
		t.Fatalf("failed to invoke Top()")
	} else {
		if top != 1 {
			t.Fatalf("got %v, but expected to be 1 after pushing 1", top)
		}
	}
}

func TestPopOpt(t *testing.T) {
	vec := []int{5, 3, 7, 11, 2}
	var h Heap
	h.InitHeap(vec)
	h.Push(1)
	h.Push(2)
	h.Pop()
	if top, _ := h.Top(); top != 2 {
		t.Fatalf("expected output to be 2, but the output is %d", top)
	}
	h.Pop()
	h.Pop()
	if top, _ := h.Top(); top != 3 {
		t.Fatalf("expected output to be 3, but the output is %d", top)
	}
}
