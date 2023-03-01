package gcoroutine

import "sync"

type SyncQueue[V any] struct {
	RQueue []V
	Mutex  sync.RWMutex
}

func NewSyncQueue[V any]() *SyncQueue[V] {
	return &SyncQueue[V]{
		RQueue: make([]V, 0),
		Mutex:  sync.RWMutex{},
	}
}

// Front 获取头部元素
func (s *SyncQueue[V]) Front() (v V) {
	s.Mutex.RLock()
	defer s.Mutex.RUnlock()
	v = s.RQueue[0]
	return
}

// Pop 弹出元素
func (s *SyncQueue[V]) Pop() (v V) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	v = s.RQueue[0]
	s.RQueue = s.RQueue[1:]
	return
}

// MultiPop 弹出多个元素
func (s *SyncQueue[V]) MultiPop(len int) (vs []V) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	for i := 0; i < len; i++ {
		vs = append(vs, s.RQueue[i])
	}
	s.RQueue = s.RQueue[len:]
	return
}

// Push 头部插入
func (s *SyncQueue[V]) Push(v V) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	s.RQueue = append([]V{v}, s.RQueue...)
}

// Empty 判断是否为空
func (s *SyncQueue[V]) Empty() bool {
	s.Mutex.RLock()
	defer s.Mutex.RUnlock()
	if len(s.RQueue) > 0 {
		return true
	}
	return false
}

// Len 长度
func (s *SyncQueue[V]) Len() int {
	s.Mutex.RLock()
	defer s.Mutex.RUnlock()
	return len(s.RQueue)
}
