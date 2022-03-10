// Package queue implements queue data structures.
package queue

import "github.com/Travis-Britz/structures/heap"

// NewPriority returns a new priority queue.
func NewPriority[T any]() *Priority[T] {
	return &Priority[T]{
		heap: heap.New(func(a item[T], b item[T]) bool {
			return a.priority > b.priority
		}),
	}
}

type item[T any] struct {
	item     T
	priority int
}

// Priority implements a priority queue using a max heap.
type Priority[T any] struct {
	heap *heap.Heap[item[T]]
}

// Push adds itm to the queue with priority.
func (q *Priority[T]) Push(itm T, priority int) {
	q.heap.Push(item[T]{
		item:     itm,
		priority: priority,
	})
}

// Pop retrieves the next item from the queue.
// If the queue was empty, more will be false and next will be the zero value of type T.
func (q *Priority[T]) Pop() (next T, more bool) {
	nxt, found := q.heap.Next()
	next = nxt.item
	return next, found
}

// Len returns the size of the queue.
func (q *Priority[T]) Len() int {
	return q.heap.Len()
}
