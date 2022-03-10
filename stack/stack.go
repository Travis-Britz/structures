package stack

// Stack implements a LIFO (last in, first out) queue.
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item ...T) {
	s.items = append(s.items, item...)
}

func (s *Stack[T]) Pop() (item T, more bool) {
	if s.Len() == 0 {
		return item, false
	}
	item = s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}

func (s *Stack[T]) Len() int {
	return len(s.items)
}
