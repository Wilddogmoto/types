package slices_test

import (
	"github.com/Wilddogmoto/types/slices"
	"testing"
)

func BenchmarkSlice_Append(b *testing.B) {

	b.Run("default_slice", func(b *testing.B) {
		b.ReportAllocs()

		//b.Log(b.N)

		slice := slices.MakeSlice[int](0, b.N)

		for i := 0; i < b.N; i++ {
			slice.Append(i)
		}

	})

	b.Run("async_slice", func(b *testing.B) {
		b.ReportAllocs()

		//b.Log(b.N)

		slice := slices.MakeSlice[int](0, b.N).Async()

		for i := 0; i < b.N; i++ {
			slice.Append(i)
		}
	})
}

func BenchmarkSlice_Sort(b *testing.B) {

	b.Run("default_sort", func(b *testing.B) {
		b.ReportAllocs()
		slice := slices.MakeSlice[int](0, b.N)

		for i := 0; i < b.N; i++ {
			slice.Append(i)
		}

		slice.Sort(nil)

	})

	b.Run("bubble_sort", func(b *testing.B) {
		b.ReportAllocs()
		slice := slices.MakeSlice[int](0, b.N)

		for i := 0; i < b.N; i++ {
			slice.Append(i)
		}

		slice.Sort(slices.BubbleSort[int])
	})

	b.Run("quick_sort", func(b *testing.B) {

		b.ReportAllocs()
		slice := slices.MakeSlice[int](0, b.N)

		for i := 0; i < b.N; i++ {
			slice.Append(i)
		}

		slice.Sort(slices.QuickSort[int])
	})

	b.Run("insertion_sort", func(b *testing.B) {
		b.ReportAllocs()
		slice := slices.MakeSlice[int](0, b.N)

		for i := 0; i < b.N; i++ {
			slice.Append(i)
		}

		slice.Sort(slices.InsertionSort[int])
	})

	b.Run("merge_sort", func(b *testing.B) {
		b.ReportAllocs()
		slice := slices.MakeSlice[int](0, b.N)

		for i := 0; i < b.N; i++ {
			slice.Append(i)
		}

		slice.Sort(slices.MergeSort[int])
	})

}
