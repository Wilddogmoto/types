package slices

import (
	"cmp"
	"sync"
)

type AsyncSlice[E cmp.Ordered] struct {
	sync.RWMutex
	slice Slice[E]
}

func (slice Slice[E]) Async() AsyncSlice[E] {
	return AsyncSlice[E]{
		RWMutex: sync.RWMutex{},
		slice:   slice,
	}
}

func (async *AsyncSlice[E]) Append(values ...E) {
	async.Lock()
	async.slice.Append(values...)
	async.Unlock()
}

func (async *AsyncSlice[E]) AppendIndex(value E, index int) {
	async.Lock()
	async.slice.AppendIndex(value, index)
	async.Unlock()
}

func (async *AsyncSlice[E]) Get(index int) E {
	async.RLock()
	defer async.RUnlock()
	return async.slice.Get(index)
}

func (async *AsyncSlice[E]) GetSlice() Slice[E] {
	async.RLock()
	defer async.RUnlock()
	return async.slice
}

func (async *AsyncSlice[E]) Copy() Slice[E] {
	async.RLock()
	defer async.RUnlock()
	return async.slice.Copy()
}

func (async *AsyncSlice[E]) Remove(index int) {
	async.Lock()
	async.slice.Remove(index)
	async.Unlock()
}

func (async *AsyncSlice[E]) Clear() {
	async.Lock()
	async.slice.Clear()
	async.Unlock()
}

func (async *AsyncSlice[E]) Cap() int {
	async.RLock()
	defer async.RUnlock()
	return async.slice.Cap()
}

func (async *AsyncSlice[E]) Iterate(fn func(int, E)) {
	async.RLock()
	async.slice.Iterate(fn)
	async.RUnlock()
}

func (async *AsyncSlice[E]) Grow(n int) {
	async.Lock()
	async.slice.Grow(n)
	async.Unlock()
}

// Sort if algorithm == nil used default algorithm sort.Sort (quicksort)
func (async *AsyncSlice[E]) Sort(algorithm SortAlgorithm[E]) {
	async.Lock()
	async.slice.Sort(algorithm)
	async.Unlock()
}

func (async *AsyncSlice[E]) Len() int {
	async.RLock()
	defer async.RUnlock()
	return async.slice.Len()
}

// Search return index and find flag (bool), if algorithm == nil, used default algorithm sort.Search (binary search)
func (async *AsyncSlice[E]) Search(value E) (int, bool) {
	async.RLock()
	defer async.RUnlock()
	return async.slice.Search(value)
}

func (async *AsyncSlice[E]) Map() map[int]E {
	async.RLock()
	defer async.RUnlock()
	return async.slice.Map()
}

func (async *AsyncSlice[E]) Max() E {
	async.RLock()
	defer async.RUnlock()
	return async.slice.Max()
}

func (async *AsyncSlice[E]) Min() E {
	async.RLock()
	defer async.RUnlock()
	return async.slice.Min()
}

func (async *AsyncSlice[E]) Compare(sl Slice[E]) int {
	async.RLock()
	defer async.RUnlock()
	return async.slice.Compare(sl)
}
