// Package goheap provides a generic heap implementation wrapping container/heap.
package goheap

import "container/heap"

// Heap is a generic heap data structure.
// By default it's a min-heap (smallest element first).
// Use NewMax for a max-heap.
type Heap[T any] struct {
	data *heapData[T]
}

type heapData[T any] struct {
	items []T
	less  func(a, b T) bool
}

func (h *heapData[T]) Len() int           { return len(h.items) }
func (h *heapData[T]) Less(i, j int) bool { return h.less(h.items[i], h.items[j]) }
func (h *heapData[T]) Swap(i, j int)      { h.items[i], h.items[j] = h.items[j], h.items[i] }

func (h *heapData[T]) Push(x any) {
	h.items = append(h.items, x.(T))
}

func (h *heapData[T]) Pop() any {
	old := h.items
	n := len(old)
	x := old[n-1]
	h.items = old[0 : n-1]
	return x
}

// New creates a min-heap with the given comparison function.
// The less function should return true if a should come before b.
// For a min-heap of integers: func(a, b int) bool { return a < b }
func New[T any](less func(a, b T) bool) *Heap[T] {
	h := &Heap[T]{
		data: &heapData[T]{
			items: make([]T, 0),
			less:  less,
		},
	}
	heap.Init(h.data)
	return h
}

// NewMax creates a max-heap with the given comparison function.
// The less function should define the natural ordering (e.g., a < b for integers).
// The heap will return the largest element first.
func NewMax[T any](less func(a, b T) bool) *Heap[T] {
	return New(func(a, b T) bool {
		return less(b, a) // reverse the comparison
	})
}

// Push adds an element to the heap.
func (h *Heap[T]) Push(x T) {
	heap.Push(h.data, x)
}

// Pop removes and returns the top element from the heap.
// Panics if the heap is empty.
func (h *Heap[T]) Pop() T {
	return heap.Pop(h.data).(T)
}

// Peek returns the top element without removing it.
// Panics if the heap is empty.
func (h *Heap[T]) Peek() T {
	return h.data.items[0]
}

// Len returns the number of elements in the heap.
func (h *Heap[T]) Len() int {
	return h.data.Len()
}

// Empty returns true if the heap has no elements.
func (h *Heap[T]) Empty() bool {
	return h.Len() == 0
}
