# heap

A generic heap (priority queue) for Go, wrapping `container/heap`.

## Installation

```bash
go get github.com/exedev/goheap
```

## Usage

### Min-Heap (default)

```go
import (
    "cmp"
    "github.com/exedev/goheap"
)

// Integer min-heap using cmp.Compare
h := goheap.New(cmp.Compare[int])
h.Push(3)
h.Push(1)
h.Push(4)

fmt.Println(h.Pop()) // 1
fmt.Println(h.Pop()) // 3
fmt.Println(h.Pop()) // 4
```

### Max-Heap

```go
// Integer max-heap
h := goheap.NewMax(cmp.Compare[int])
h.Push(3)
h.Push(1)
h.Push(4)

fmt.Println(h.Pop()) // 4
fmt.Println(h.Pop()) // 3
fmt.Println(h.Pop()) // 1
```

### Custom Structs

```go
type Task struct {
    Priority int
    Name     string
}

// Using a cmp function (returns negative, zero, or positive)
h := goheap.New(func(a, b Task) int {
    return a.Priority - b.Priority
})

h.Push(Task{3, "low"})
h.Push(Task{1, "high"})
h.Push(Task{2, "medium"})

fmt.Println(h.Pop().Name) // high
fmt.Println(h.Pop().Name) // medium
fmt.Println(h.Pop().Name) // low
```

### Using a Less Function

```go
// NewMin accepts a less function (returns bool)
h := goheap.NewMin(func(a, b int) bool { return a < b })
```

## API

- `New[T](cmp func(a, b T) int) *Heap[T]` - Create a min-heap with a cmp function
- `NewMax[T](cmp func(a, b T) int) *Heap[T]` - Create a max-heap with a cmp function
- `NewMin[T](less func(a, b T) bool) *Heap[T]` - Create a min-heap with a less function
- `Push(x T)` - Add an element
- `Pop() T` - Remove and return the top element
- `Peek() T` - Return the top element without removing it
- `Len() int` - Number of elements
- `Empty() bool` - Check if heap is empty
