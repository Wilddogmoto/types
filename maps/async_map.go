package maps

import (
	"cmp"
	"sync"
)

type (
	AsyncMap[T Map[Key, Val], Key cmp.Ordered, Val any] struct {
		sync.RWMutex
		storage Map[Key, Val]
	}
)

func MakeAsyncMap[T Map[Key, Val], Key cmp.Ordered, Val any](storage Map[Key, Val]) AsyncMap[T, Key, Val] {
	return AsyncMap[T[Key, Val], Key, Val]{
		RWMutex: sync.RWMutex{},
		storage: storage,
	}
}

func (m *AsyncMap[T, Key, Val]) Add(key Key, value Val) {
	m.Lock()
	m.storage.Add(key, value)
	m.Unlock()
}

func (m *AsyncMap[T, Key, Val]) Delete(key Key) {
	m.Lock()
	m.storage.Delete(key)
	m.Unlock()
}

func (m *AsyncMap[T, Key, Val]) Contains(key Key) bool {
	m.RLock()
	defer m.RUnlock()
	return m.storage.Contains(key)
}

func (m *AsyncMap[T, Key, Val]) Get(key Key) (Val, bool) {
	m.RLock()
	defer m.RUnlock()
	return m.storage.Get(key)
}

func (m *AsyncMap[T, Key, Val]) Iterate(fn func(Key, Val)) {
	m.RLock()
	defer m.RUnlock()
	m.storage.Iterate(fn)
}

func (m *AsyncMap[T, Key, Val]) Len() int {
	m.RLock()
	defer m.RUnlock()
	return m.storage.Len()
}

func (m *AsyncMap[T, Key, Val]) Copy() Map[Key, Val] {
	m.RLock()
	defer m.RUnlock()
	return m.storage.Copy()
}

func (m *AsyncMap[T, Key, Val]) Clear() {
	m.RLock()
	m.storage.Clear()
	defer m.RUnlock()
}
