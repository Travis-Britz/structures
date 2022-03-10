package queue_test

import (
	"fmt"
	"testing"

	"github.com/Travis-Britz/structures/queue"
)

func ExampleNewPriority() {
	q := queue.NewPriority[string]()
	q.Push("a", 1)
	q.Push("b", 5)
	q.Push("c", 10)
	fmt.Println("Len:", q.Len())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	// Output:
	// Len: 3
	// c true
	// b true
	// a true
	//  false
	//  false

}

func TestNewPriority(t *testing.T) {
	q := queue.NewPriority[int]()
	q.Push(1, -1)
	q.Push(2, 1)
	q.Push(3, 5)

	for _, expected := range []int{3, 2, 1, 0} {
		next, _ := q.Pop()
		if next != expected {
			t.Errorf("expected %v; got %v", expected, next)
		}
	}
	if _, more := q.Pop(); more {
		t.Errorf("expected queue to be empty")
	}
}
