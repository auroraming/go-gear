package gcoroutine

import "sync"

type SyncSlice[V any] struct {
	RSlice []V
	Mutex  sync.RWMutex
}

func NewSyncSlice[V any]() *SyncSlice[V] {
	return &SyncSlice[V]{
		RSlice: make([]V, 0),
		Mutex:  sync.RWMutex{},
	}
}

// Load 增加
func (s *SyncSlice[V]) Load(pos int) (value V) {
	s.Mutex.RLock()
	defer s.Mutex.RUnlock()
	value = s.RSlice[pos]
	return
}

func (s *SyncSlice[V]) Store(pos int, value V) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	s.RSlice[pos] = value
}

func (s *SyncSlice[V]) Append(value V) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	s.RSlice = append(s.RSlice, value)
}

// Delete 删除
func (s *SyncSlice[V]) Delete(pos int) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	s.RSlice = append(s.RSlice[0:pos], s.RSlice[pos+1:]...)
}

// Range 查询
func (s *SyncSlice[V]) Range(fn func(pos int, value V) bool) {
	s.Mutex.RLock()
	defer s.Mutex.RUnlock()
	for pos, value := range s.RSlice {
		if !fn(pos, value) {
			break
		}
	}
}

func (s *SyncSlice[V]) Len() int {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	return len(s.RSlice)
}

// XRange 批量获取
func (s *SyncSlice[V]) XRange(start, end int) []V {
	s.Mutex.RLock()
	defer s.Mutex.RUnlock()
	if end > len(s.RSlice)-1 || start < 0 {
		return nil
	}
	arr := make([]V, end-start+1)
	copy(arr, s.RSlice[start:end+1])
	return arr
}

// Remove 删除
func (s *SyncSlice[V]) Remove(pos int) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	s.RSlice = append(s.RSlice[:pos], s.RSlice[pos+1:]...)
}

// XRemove 删除
func (s *SyncSlice[V]) XRemove(start, end int) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	if start < 0 {
		return
	}
	if end > len(s.RSlice)-1 && start == 0 {
		s.RSlice = []V{}
	} else if end > len(s.RSlice)-1 {
		end = len(s.RSlice) - 1
		s.RSlice = append(s.RSlice[:start], s.RSlice[end+1:]...)
	} else {
		s.RSlice = append(s.RSlice[:start], s.RSlice[end+1:]...)
	}
}
