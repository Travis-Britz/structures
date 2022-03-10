package stack_test

import (
	"fmt"

	"github.com/Travis-Britz/structures/stack"
)

func ExampleStack_Push() {
	s := &stack.Stack[int]{}
	s.Push(1)
	s.Push(2, 3, 4)
	_, _ = s.Pop()
	s.Push()
	s.Push(5)
	for next, more := s.Pop(); more; {
		fmt.Println(next, more)
	}
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	// Output:
	// 5 true
	// 3 true
	// 2 true
	// 1 true
	// 0 false
	// 0 false
	// 0 false
}
