package heap_test

import (
	"testing"

	"github.com/Travis-Britz/structures/heap"
)

func TestHeap_MinHeap(t *testing.T) {

	tt := []struct {
		given    []int
		expected []int
	}{{
		given:    []int{2, 4, 5, 1, 3},
		expected: []int{1, 2, 3, 4, 5},
	}, {
		given:    []int{1, 3, 2, 4},
		expected: []int{1, 2, 3, 4},
	}, {
		given:    []int{5, 4, 3, 2, 1},
		expected: []int{1, 2, 3, 4, 5},
	}, {
		given:    []int{1, 2, 3, 4, 5},
		expected: []int{1, 2, 3, 4, 5},
	}, {
		given:    []int{6, 5, 4, 3, 2, 1},
		expected: []int{1, 2, 3, 4, 5, 6},
	}, {
		given:    []int{7, 6, 5, 4, 3, 2, 1},
		expected: []int{1, 2, 3, 4, 5, 6, 7},
	}, {
		given:    []int{5, 5, 4, 3, 1, 1},
		expected: []int{1, 1, 3, 4, 5, 5},
	},
	}

	// smallest should come out first
	compareFn := heap.Min[int]

	for _, tt := range tt {
		t.Run("", func(t *testing.T) {
			q := heap.New(compareFn)
			for _, item := range tt.given {
				q.Push(item)
			}
			for _, expected := range tt.expected {
				next, more := q.Next()
				if !more {
					break
				}
				if next != expected {
					t.Errorf("expected %v; got %v", expected, next)
				}
			}
		})
	}
}

func TestHeap_MaxHeap(t *testing.T) {

	tt := []struct {
		given    []int
		expected []int
	}{{
		given:    []int{2, 4, 5, 1, 3},
		expected: []int{5, 4, 3, 2, 1},
	}, {
		given:    []int{1, 3, 2, 4},
		expected: []int{4, 3, 2, 1},
	}, {
		given:    []int{5, 4, 3, 2, 1},
		expected: []int{5, 4, 3, 2, 1},
	}, {
		given:    []int{1, 2, 3, 4, 5},
		expected: []int{5, 4, 3, 2, 1},
	}, {
		given:    []int{6, 5, 4, 3, 2, 1},
		expected: []int{6, 5, 4, 3, 2, 1},
	}, {
		given:    []int{7, 6, 5, 4, 3, 2, 1},
		expected: []int{7, 6, 5, 4, 3, 2, 1},
	}, {
		given:    []int{5, 5, 4, 3, 1, 1},
		expected: []int{5, 5, 4, 3, 1, 1},
	},
	}

	// smallest should come out first
	compareFn := heap.Max[int]

	for _, tt := range tt {
		t.Run("", func(t *testing.T) {
			q := heap.New(compareFn)
			for _, item := range tt.given {
				q.Push(item)
			}
			for _, expected := range tt.expected {
				next, more := q.Next()
				if !more {
					break
				}
				if next != expected {
					t.Errorf("expected %v; got %v", expected, next)
				}
			}
		})
	}
}

func TestHeap_NewUsing(t *testing.T) {

	tt := []struct {
		given    []int
		expected []int
	}{{
		given:    []int{2, 4, 5, 1, 3},
		expected: []int{1, 2, 3, 4, 5},
	}, {
		given:    []int{1, 3, 2, 4},
		expected: []int{1, 2, 3, 4},
	}, {
		given:    []int{5, 4, 3, 2, 1},
		expected: []int{1, 2, 3, 4, 5},
	}, {
		given:    []int{1, 2, 3, 4, 5},
		expected: []int{1, 2, 3, 4, 5},
	}, {
		given:    []int{6, 5, 4, 3, 2, 1},
		expected: []int{1, 2, 3, 4, 5, 6},
	}, {
		given:    []int{7, 6, 5, 4, 3, 2, 1},
		expected: []int{1, 2, 3, 4, 5, 6, 7},
	}, {
		given:    []int{5, 5, 4, 3, 1, 1},
		expected: []int{1, 1, 3, 4, 5, 5},
	},
	}

	for _, td := range tt {
		t.Run("", func(t *testing.T) {
			given := make([]int, len(td.given))
			copy(given, td.given)
			q := heap.NewUsing(given, heap.Min[int])
			for _, expected := range td.expected {
				next, more := q.Next()
				if !more {
					break
				}
				if next != expected {
					t.Errorf("expected %v; got %v", expected, next)
				}
			}
		})
	}
}

func TestHeapEmpty(t *testing.T) {
	q := heap.New(heap.Min[int])
	next, more := q.Next()
	if next != 0 {
		t.Errorf("expected zero value for next; got %v", next)
	}
	if more != false {
		t.Errorf("expected more to be false on an empty queue")
	}
}

func TestHeapVariadicPush(t *testing.T) {
	q := heap.New(heap.Min[int])
	q.Push(2, 1, 3)
	for _, expected := range []int{1, 2, 3} {
		next, _ := q.Next()
		if expected != next {
			t.Errorf("expected %v; got %v", expected, next)
		}
	}
	q.Push()
	_, more := q.Next()
	if more != false {
		t.Errorf("expected queue to be empty")
	}
}
