# goheap

A generic heap (priority queue) for Go, wrapping `container/heap`.

## Installation

```bash
go get github.com/exedev/goheap
```

## Usage

### Min-Heap (default)

```go
import "github.com/exedev/goheap"

// Integer min-heap
h := goheap.New(func(a, b int) bool { return a < b })
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
h := goheap.NewMax(func(a, b int) bool { return a < b })
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

h := goheap.New(func(a, b Task) bool {
    return a.Priority < b.Priority
})

h.Push(Task{3, "low"})
h.Push(Task{1, "high"})
h.Push(Task{2, "medium"})

fmt.Println(h.Pop().Name) // high
fmt.Println(h.Pop().Name) // medium
fmt.Println(h.Pop().Name) // low
```

## API

- `New[T](less func(a, b T) bool) *Heap[T]` - Create a min-heap
- `NewMax[T](less func(a, b T) bool) *Heap[T]` - Create a max-heap
- `Push(x T)` - Add an element
- `Pop() T` - Remove and return the top element
- `Peek() T` - Return the top element without removing it
- `Len() int` - Number of elements
- `Empty() bool` - Check if heap is empty
