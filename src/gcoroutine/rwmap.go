package gcoroutine

import (
	"sync"
)

type SyncMap[K string | uintptr | uint8 | int | int32 | int64 | uint | uint16 | uint32 | uint64, V any] struct {
	RMap  map[K]V
	Mutex *sync.RWMutex
}

func NewSyncMap[K string | uintptr | uint8 | uint16 | uint | uint32 | uint64 | int | int32 | int64, V any]() *SyncMap[K, V] {
	return &SyncMap[K, V]{
		RMap:  make(map[K]V),
		Mutex: new(sync.RWMutex),
	}
}

// Load 增加
func (m *SyncMap[K, V]) Load(key K) (value V, ok bool) {
	m.Mutex.RLock()
	defer m.Mutex.RUnlock()
	value, ok = m.RMap[key]
	return
}

// LoadOrStore 不存就添加
func (m *SyncMap[K, V]) LoadOrStore(key K, value V) {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()
	_, ok := m.RMap[key]
	if !ok {
		m.RMap[key] = value
	}
}

// Store 存储
func (m *SyncMap[K, V]) Store(key K, value V) {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()
	m.RMap[key] = value
}

// Delete 删除
func (m *SyncMap[K, V]) Delete(key K) {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()
	delete(m.RMap, key)
}

// Range 查询
func (m *SyncMap[K, V]) Range(fn func(key K, value V) bool) {
	m.Mutex.RLock()
	defer m.Mutex.RUnlock()
	for key, value := range m.RMap {
		if !fn(key, value) {
			break
		}
	}

}

// Len 获取长度
func (m *SyncMap[K, V]) Len() int {
	m.Mutex.RLock()
	defer m.Mutex.RUnlock()
	return len(m.RMap)
}

// Copy 拷贝
func (m *SyncMap[K, V]) Copy(ma map[K]V) {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()
	for k, v := range m.RMap {
		ma[k] = v
	}
}
