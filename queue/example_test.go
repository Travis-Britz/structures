package queue_test

import (
	"fmt"

	"github.com/Travis-Britz/structures/queue"
)

func ExampleNew_int() {
	vals := []int{2, 3, 1, 5, 4}
	q := queue.New(queue.Min[int])
	for _, v := range vals {
		q.Push(v)
	}
	for next, more := q.Next(); more; {
		fmt.Println(next)
	}
	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
}

func ExampleNew_string() {
	vals := []string{"a", "b", "c", "d", "e"}
	q := queue.New(queue.Max[string])
	for _, v := range vals {
		q.Push(v)
	}
	for next, more := q.Next(); more; {
		fmt.Println(next)
	}
	// Output:
	// e
	// d
	// c
	// b
	// a
}

// This example demonstrates creating a queue with a custom type and comparison function.
func ExampleNew_custom() {
	type myType struct {
		complexField struct{}
		priority     int
	}
	compareFn := func(a myType, b myType) bool {
		return a.priority > b.priority
	}
	vals := []myType{{priority: 1}, {priority: 4}, {priority: 2}}
	q := queue.New(compareFn)
	for _, v := range vals {
		q.Push(v)
	}
	for next, more := q.Next(); more; {
		fmt.Println(next.priority)
	}
	// Output:
	// 4
	// 2
	// 1

}

func ExampleNewUsing_int() {
	vals := []int{3, 5, 2, 1, 4}
	q := queue.NewUsing(vals, queue.Min[int])
	for next, more := q.Next(); more; {
		fmt.Println(next)
	}
	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
}

func ExampleNewUsing_capacity() {
	vals := []int{3, 5, 2, 1, 4}
	// the backing array for the heap will start out with a capacity of 10000
	q := queue.NewUsing(make([]int, 0, 10000), queue.Min[int])
	for _, v := range vals {
		q.Push(v)
	}
	for next, more := q.Next(); more; {
		fmt.Println(next)
	}
	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
}

func ExampleMax() {
	q := queue.New(queue.Max[int])
	for _, v := range []int{1, 3, 2} {
		q.Push(v)
	}
	for next, more := q.Next(); more; {
		fmt.Println(next)
	}
	// Output:
	// 3
	// 2
	// 1
}
