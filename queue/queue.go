// Package queue implements a priority queue data structure using a binary heap.
package queue

import (
	"golang.org/x/exp/constraints"
)

// New returns a new Priority queue using the comparison function fn.
func New[T any](fn func(T, T) bool) *Priority[T] {
	return &Priority[T]{compare: fn}
}

// NewUsing returns a new queue with the given slice as the underlying array.
// Callers must not access slice again.
// A zero-length slice with a high capacity may be provided to prevent Push from growing the array.
// If the slice contains any elements, it will be heapified according to fn.
func NewUsing[T any](slice []T, fn func(T, T) bool) *Priority[T] {
	q := &Priority[T]{
		items:   slice,
		compare: fn,
	}
	q.heapify()
	return q
}

type Priority[T any] struct {
	items   []T
	compare func(T, T) bool
}

func (q *Priority[T]) Push(i T) {
	q.items = append(q.items, i)
	q.percUp(len(q.items) - 1)
}

func (q *Priority[T]) Next() (next T, more bool) {
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

func (q *Priority[T]) percUp(i int) {
	pi := (i - 1) / 2 // parent index
	if q.compare(q.items[i], q.items[pi]) {
		q.items[i], q.items[pi] = q.items[pi], q.items[i]
		q.percUp(pi)
	}
}

func (q *Priority[T]) percDown(i int) {
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

func (q *Priority[T]) heapify() {
	for i := len(q.items) - 1; i >= 0; i-- {
		q.percDown(i)
	}
}

// Max is a comparison function that turns the queue into a Max Heap.
func Max[T constraints.Ordered](before T, after T) bool {
	return before > after
}

// Min is a comparison function that turns the queue into a Min Heap.
func Min[T constraints.Ordered](before T, after T) bool {
	return before < after
}
