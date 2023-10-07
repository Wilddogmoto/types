package maps

import (
	"cmp"
	"sync"
)

type (
	AsyncMap[Key cmp.Ordered, Val any] struct {
		sync.RWMutex
		storage Map[Key, Val]
	}
)

func (m Map[Key, Val]) Async() *AsyncMap[Key, Val] {
	return &AsyncMap[Key, Val]{
		RWMutex: sync.RWMutex{},
		storage: m,
	}
}

func (m *AsyncMap[Key, Val]) Add(key Key, value Val) {
	m.Lock()
	m.storage.Add(key, value)
	m.Unlock()
}

func (m *AsyncMap[Key, Val]) Delete(key Key) {
	m.Lock()
	m.storage.Delete(key)
	m.Unlock()
}

func (m *AsyncMap[Key, Val]) Contains(key Key) bool {
	m.RLock()
	defer m.RUnlock()
	return m.storage.Contains(key)
}

func (m *AsyncMap[Key, Val]) Get(key Key) (Val, bool) {
	m.RLock()
	defer m.RUnlock()
	return m.storage.Get(key)
}

func (m *AsyncMap[Key, Val]) Iterate(fn func(Key, Val)) {
	m.RLock()
	defer m.RUnlock()
	m.storage.Iterate(fn)
}

func (m *AsyncMap[Key, Val]) Len() int {
	m.RLock()
	defer m.RUnlock()
	return m.storage.Len()
}

func (m *AsyncMap[Key, Val]) Copy() Map[Key, Val] {
	m.RLock()
	defer m.RUnlock()
	return m.storage.Copy()
}

func (m *AsyncMap[Key, Val]) Clear() {
	m.RLock()
	m.storage.Clear()
	defer m.RUnlock()
}

func (m *AsyncMap[Key, Val]) Values() []Val {
	m.RLock()
	defer m.RUnlock()
	return m.storage.Values()
}

func (m *AsyncMap[Key, Val]) Keys() []Key {

	m.RLock()
	defer m.RUnlock()
	return m.storage.Keys()
}

func (m *AsyncMap[Key, Val]) Merge(input Map[Key, Val]) {
	m.Lock()
	m.storage.Merge(input)
	m.Unlock()
}
