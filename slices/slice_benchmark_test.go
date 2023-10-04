package slices_test

import (
	"github.com/Wilddogmoto/types/slices"
	"testing"
)

func BenchmarkSlice_Append(b *testing.B) {
	b.ReportAllocs()
	slice := slices.MakeSlice[int](0, b.N)

	for i := 0; i < b.N; i++ {
		slice.Append(i)
	}
}

func BenchmarkAsyncSlice_Append(b *testing.B) {
	b.ReportAllocs()
	slice := slices.MakeSlice[int](0, b.N).Async()

	for i := 0; i < b.N; i++ {
		slice.Append(i)
	}
}
