package gcoroutine

import "testing"

func TestQueue(t *testing.T) {
	queue := NewSyncQueue[int]()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
}
