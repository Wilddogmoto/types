package maps

import (
	"cmp"
	"maps"
)

type (
	Map[Key cmp.Ordered, Val any] map[Key]Val
)

func MakeMap[Key cmp.Ordered, Val any](size int) Map[Key, Val] {
	return make(Map[Key, Val], size)
}

func (m Map[Key, Val]) Add(key Key, value Val) {
	m[key] = value
}

func (m Map[Key, Val]) Delete(key Key) {
	delete(m, key)
}

func (m Map[Key, Val]) Contains(key Key) bool {
	_, ok := m[key]
	return ok
}

func (m Map[Key, Val]) Get(key Key) (Val, bool) {

	val, ok := m[key]

	return val, ok
}

func (m Map[Key, Val]) Iterate(f func(key Key, value Val)) {
	for key, value := range m {
		f(key, value)
	}
}

func (m Map[Key, Val]) Len() int {
	return len(m)
}

func (m Map[Key, Val]) Copy() Map[Key, Val] {
	newMap := make(Map[Key, Val], m.Len())
	maps.Copy(newMap, m)
	return newMap
}

func (m *Map[Key, Val]) Clear() {
	//maps.Clear(m) - if there is a large map, it will take a long time to clean up
	*m = make(Map[Key, Val], m.Len())
}

func (m Map[Key, Val]) Values() []Val {

	out := make([]Val, 0, m.Len())

	for _, value := range m {
		out = append(out, value)
	}

	return out
}

func (m Map[Key, Val]) Keys() []Key {

	out := make([]Key, 0, m.Len())

	for key := range m {
		out = append(out, key)
	}

	return out
}

func (m Map[Key, Val]) Merge(input Map[Key, Val]) {
	maps.Copy(m, input)
}
