// Package heap implements a binary heap.
package heap

// New returns a new heap using the comparison function fn.
func New[T any](fn func(T, T) bool) *Heap[T] {
	return &Heap[T]{compare: fn}
}

// NewUsing returns a new heap with the given slice as the underlying array.
// Callers must not access slice again.
// A zero-length slice with a high capacity may be provided to prevent Push from growing the array.
// If the slice contains any elements, it will be heapified according to fn.
func NewUsing[T any](slice []T, fn func(T, T) bool) *Heap[T] {
	q := &Heap[T]{
		items:   slice,
		compare: fn,
	}
	q.heapify()
	return q
}

// Heap implements a binary heap, sorted according to a comparison function.
type Heap[T any] struct {
	items   []T
	compare func(T, T) bool
}

// Push adds an item (or items) to the heap.
func (q *Heap[T]) Push(item ...T) {
	if len(item) != 1 {
		for _, v := range item {
			q.Push(v)
		}
		return
	}
	q.items = append(q.items, item[0])
	q.percUp(len(q.items) - 1)
}

// Next removes the next item from the top of the heap.
// If the heap was empty, more will be false and the next item will be the zero value of type T.
func (q *Heap[T]) Next() (next T, more bool) {
	if len(q.items) == 0 {
		var zero T
		return zero, false
	}

	result := q.items[0]
	q.items[0] = q.items[len(q.items)-1]
	q.items = q.items[:len(q.items)-1]
	q.percDown(0)
	return result, true
}

// Len returns the length of the queue.
func (q *Heap[T]) Len() int {
	return len(q.items)
}

func (q *Heap[T]) percUp(i int) {
	pi := (i - 1) / 2 // parent index
	if q.compare(q.items[i], q.items[pi]) {
		q.items[i], q.items[pi] = q.items[pi], q.items[i]
		q.percUp(pi)
	}
}

func (q *Heap[T]) percDown(i int) {
	left, right := i*2+1, i*2+2
	last := len(q.items) - 1

	switch {
	case left > last:
		return
	case right > last:
		if q.compare(q.items[left], q.items[i]) {
			q.items[i], q.items[left] = q.items[left], q.items[i]
		}
	case q.compare(q.items[left], q.items[right]):
		if q.compare(q.items[left], q.items[i]) {
			q.items[i], q.items[left] = q.items[left], q.items[i]
			q.percDown(left)
		}
	default:
		if q.compare(q.items[right], q.items[i]) {
			q.items[i], q.items[right] = q.items[right], q.items[i]
			q.percDown(right)
		}
	}
}

func (q *Heap[T]) heapify() {
	for i := (len(q.items) - 1) / 2; i >= 0; i-- {
		q.percDown(i)
	}
}

// Max is a comparison function that turns the queue into a Max Heap.
func Max[T ordered](before T, after T) bool {
	return before > after
}

// Min is a comparison function that turns the queue into a Min Heap.
func Min[T ordered](before T, after T) bool {
	return before < after
}

type ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
	~float32 | ~float64 |
	~string
}
