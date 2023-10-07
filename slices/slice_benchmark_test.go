package slices_test

import (
	"fmt"
	"github.com/Wilddogmoto/types/slices"
	"math/rand"
	"testing"
)

func newSlice(size int) (slice slices.Slice[int]) {

	slice = slices.MakeSlice[int](0, size)

	for i := 0; i < size; i++ {
		slice.Append(rand.Intn(size))
	}
	return slice
}

func BenchmarkSlice_Append(b *testing.B) {

	b.Run("default_slice", func(b *testing.B) {
		b.ReportAllocs()

		slice := slices.MakeSlice[int](0, b.N)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			slice.Append(i)
		}

	})

	b.Run("async_slice", func(b *testing.B) {
		b.ReportAllocs()

		slice := slices.MakeSlice[int](0, b.N).Async()

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			slice.Append(i)
		}
	})
}

func BenchmarkSlice_Sort(b *testing.B) {

	tests := []struct {
		name      string
		algorithm slices.SortAlgorithm[int]
	}{
		{
			name:      "default_sort",
			algorithm: nil,
		},
		{
			name:      "BubbleSort",
			algorithm: slices.BubbleSort[int],
		},
		{
			name:      "InsertionSort",
			algorithm: slices.InsertionSort[int],
		},
		{
			name:      "MergeSort",
			algorithm: slices.MergeSort[int],
		},
		{
			name:      "QuickSort",
			algorithm: slices.QuickSort[int],
		},
		{
			name:      "HeapSort",
			algorithm: slices.HeapSort[int],
		},
	}

	benchs := []struct {
		size  int
		tests []struct {
			name      string
			algorithm slices.SortAlgorithm[int]
		}
	}{
		{
			size:  100,
			tests: tests,
		},
		{
			size:  1000,
			tests: tests,
		},
		{
			size:  10000,
			tests: tests,
		},
		{
			size:  100000,
			tests: tests,
		},
	}

	for _, bench := range benchs {

		slice := newSlice(bench.size)

		for _, test := range bench.tests {
			b.Run(fmt.Sprintf("%s_%d", test.name, bench.size), func(b *testing.B) {

				b.ReportAllocs()

				newsl := slice.Copy()

				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					newsl.Sort(test.algorithm)
				}
			})
		}
	}
}
