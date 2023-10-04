package slices

import (
	"cmp"
	"slices"
)

type Slice[E cmp.Ordered] []E

// MakeSlice cap should not be less len
func MakeSlice[E cmp.Ordered](len, cap int) Slice[E] {
	return make(Slice[E], len, cap)
}

func (slice *Slice[E]) Append(values ...E) {
	*slice = append(*slice, values...)
}

func (slice Slice[E]) AppendIndex(value E, index int) {
	slice[index] = value
}

func (slice Slice[E]) Get(index int) E {
	return slice[index]
}

func (slice Slice[E]) Copy() Slice[E] {
	out := make(Slice[E], slice.Len())
	copy(out, slice)
	return out
}

func (slice *Slice[E]) Remove(index int) {
	ss := *slice
	copy(ss[index:], ss[index+1:])
	*slice = ss[:ss.Len()-1]
}

func (slice *Slice[E]) Clear() {
	ss := *slice
	*slice = ss[:0]
}

func (slice Slice[E]) Cap() int { return cap(slice) }

func (slice Slice[E]) Iterate(f func(index int, value E)) {

	for index, value := range slice {
		f(index, value)
	}
}

func (slice *Slice[E]) Grow(n int) {

	if n < 0 {
		panic("cannot be negative")
	}
	if n -= slice.Cap() - slice.Len(); n > 0 {
		sl := *slice
		*slice = append(sl[:slice.Cap()], make(Slice[E], n)...)[:slice.Len()]
	}
}

// Sort if algorithm == nil used default algorithm sort.Sort (quicksort).
// Внимание русские символы по типу `Ё` могут не проходить сортировку
func (slice *Slice[E]) Sort(algorithm SortAlgorithm[E]) {
	if algorithm != nil {
		*slice = algorithm(*slice)
		return
	}

	slices.SortFunc(*slice, func(a, b E) int {
		return cmp.Compare(a, b)
	})
}

func (slice Slice[E]) Len() int { return len(slice) }

// Search return index and find flag (bool), used default algorithm slices.BinarySearch (binary search)
// Attention, the slice needs to be sorted before searching
func (slice Slice[E]) Search(value E) (int, bool) {

	lenSlice := slice.Len()

	if lenSlice == 0 {
		return 0, false
	}

	return slices.BinarySearch(slice, value)
}

func (slice Slice[E]) Map() map[int]E {

	out := make(map[int]E, slice.Len())

	for index, value := range slice {
		out[index] = value
	}

	return out
}

func (slice Slice[E]) Max() E { return slices.Max(slice) }

func (slice Slice[E]) Min() E { return slices.Min(slice) }

func (slice Slice[E]) Compare(sl Slice[E]) int {
	return slices.Compare(slice, sl)
}
