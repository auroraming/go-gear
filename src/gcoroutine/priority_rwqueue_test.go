package gcoroutine

import (
	"fmt"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	queue := NewSyncPriorityQueue[int]()
	queue.Push(100, 1)
	queue.Push(1, 2)
	queue.Push(2, 3)
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
}
