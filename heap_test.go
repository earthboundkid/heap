package goheap

import (
	"cmp"
	"testing"
)

func TestMinHeapInt(t *testing.T) {
	h := New(cmp.Compare[int])

	h.Push(3)
	h.Push(1)
	h.Push(4)
	h.Push(1)
	h.Push(5)

	if h.Len() != 5 {
		t.Errorf("expected len 5, got %d", h.Len())
	}

	expected := []int{1, 1, 3, 4, 5}
	for i, want := range expected {
		if h.Empty() {
			t.Fatalf("heap empty at index %d", i)
		}
		got := h.Pop()
		if got != want {
			t.Errorf("pop %d: expected %d, got %d", i, want, got)
		}
	}

	if !h.Empty() {
		t.Error("expected heap to be empty")
	}
}

func TestMaxHeapInt(t *testing.T) {
	h := NewMax(cmp.Compare[int])

	h.Push(3)
	h.Push(1)
	h.Push(4)
	h.Push(1)
	h.Push(5)

	expected := []int{5, 4, 3, 1, 1}
	for i, want := range expected {
		got := h.Pop()
		if got != want {
			t.Errorf("pop %d: expected %d, got %d", i, want, got)
		}
	}
}

func TestPeek(t *testing.T) {
	h := New(cmp.Compare[int])

	h.Push(5)
	h.Push(3)
	h.Push(7)

	if h.Peek() != 3 {
		t.Errorf("expected peek 3, got %d", h.Peek())
	}

	// Peek shouldn't remove the element
	if h.Len() != 3 {
		t.Errorf("expected len 3 after peek, got %d", h.Len())
	}
}

type task struct {
	priority int
	name     string
}

func TestNewMinWithLess(t *testing.T) {
	h := NewMin(func(a, b int) bool { return a < b })

	h.Push(3)
	h.Push(1)
	h.Push(4)

	expected := []int{1, 3, 4}
	for i, want := range expected {
		got := h.Pop()
		if got != want {
			t.Errorf("pop %d: expected %d, got %d", i, want, got)
		}
	}
}

func TestStructHeap(t *testing.T) {
	h := New(func(a, b task) int { return a.priority - b.priority })

	h.Push(task{3, "low"})
	h.Push(task{1, "high"})
	h.Push(task{2, "medium"})

	expected := []string{"high", "medium", "low"}
	for i, want := range expected {
		got := h.Pop()
		if got.name != want {
			t.Errorf("pop %d: expected %s, got %s", i, want, got.name)
		}
	}
}

func TestStringHeap(t *testing.T) {
	h := New(cmp.Compare[string])

	h.Push("banana")
	h.Push("apple")
	h.Push("cherry")

	expected := []string{"apple", "banana", "cherry"}
	for i, want := range expected {
		got := h.Pop()
		if got != want {
			t.Errorf("pop %d: expected %s, got %s", i, want, got)
		}
	}
}
