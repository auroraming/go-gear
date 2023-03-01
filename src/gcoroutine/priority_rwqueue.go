package gcoroutine

import (
	"sync"
)

// 优先级队列
// 两种方案
// 1. Map + Slice
// 2. 双链表结构
type Value[V any] struct {
	pre      *Value[V] // 前结点
	priority int       // 优先级
	v        V         // 内容
	next     *Value[V] // 后节点
}

type SyncPriorityQueue[V any] struct {
	head  *Value[V]
	len   int
	Mutex sync.RWMutex
}

func NewSyncPriorityQueue[V any]() *SyncPriorityQueue[V] {
	return &SyncPriorityQueue[V]{
		head:  nil,
		len:   0,
		Mutex: sync.RWMutex{},
	}
}

// Front 获取头部元素
func (s *SyncPriorityQueue[V]) Front() (v V) {
	s.Mutex.RLock()
	defer s.Mutex.RUnlock()
	if s.head == nil {
		return
	}
	v = s.head.v
	return
}

// Pop 弹出元素
func (s *SyncPriorityQueue[V]) Pop() (v V) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	if s.head == nil {
		return
	}
	v = s.head.v
	if s.head.next == nil {
		s.head = nil
	} else {
		s.head = s.head.next
	}
	s.len--
	return
}

// MultiPop 弹出多个元素
func (s *SyncPriorityQueue[V]) MultiPop(len int) (vs []V) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	head := s.head
	if head == nil {
		return
	}
	for i := 0; i < len; i++ {
		vs = append(vs, head.v)
		head = head.next
		s.len--
		if head == nil {
			return
		}
	}
	s.head = head
	return
}

// Push 头部插入 priority 优先级
func (s *SyncPriorityQueue[V]) Push(priority int, v V) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	head := s.head
	node := &Value[V]{
		priority: priority,
		v:        v,
	}
	if head == nil {
		s.head = node
		s.len++
		return
	}
	for head != nil {
		if priority > head.priority {
			node.pre = head.pre
			node.next = head
			head.pre.next = node
			head.pre = node
			s.len++
			return
		}
		if head.next == nil {
			// 在最后
			head.next = node
			node.pre = head
			s.len++
			return
		}
		head = head.next
	}
}

// Empty 判断是否为空
func (s *SyncPriorityQueue[V]) Empty() bool {
	s.Mutex.RLock()
	defer s.Mutex.RUnlock()
	if s.len > 0 {
		return true
	}
	return false
}

// Len 长度
func (s *SyncPriorityQueue[V]) Len() int {
	s.Mutex.RLock()
	defer s.Mutex.RUnlock()
	return s.len
}

// Clear 清空
func (s *SyncPriorityQueue[V]) Clear() {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	s.head = nil
	s.len = 0
}
