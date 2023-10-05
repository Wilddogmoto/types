package slices_test

import (
	"cmp"
	"github.com/Wilddogmoto/types/slices"
	"testing"
)

func TestSortInteger(t *testing.T) {

	add := []int{3, 2, 33, 100, 1, 0, 0, 0}
	want := []int{0, 0, 0, 1, 2, 3, 33, 100}

	tests := []struct {
		name      string
		add       slices.Slice[int]
		want      slices.Slice[int]
		algorithm slices.SortAlgorithm[int]
	}{
		{
			name:      "default sort",
			add:       add,
			want:      want,
			algorithm: nil,
		},
		{
			name:      "bubble sort",
			add:       add,
			want:      want,
			algorithm: slices.BubbleSort[int],
		},
		{
			name:      "insertion sort",
			add:       add,
			want:      want,
			algorithm: slices.InsertionSort[int],
		},
		{
			name:      "merge sort",
			add:       add,
			want:      want,
			algorithm: slices.MergeSort[int],
		},
		{
			name:      "quick sort",
			add:       add,
			want:      want,
			algorithm: slices.QuickSort[int],
		},
		{
			name: "any sort",
			add:  add,
			want: want,
			algorithm: slices.AnySort[int](func(a, b int) int {
				return cmp.Compare(a, b)
			}),
		},
	}

	slice := slices.MakeSlice[int](0, 0)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			slice.Append(tt.add...)
			defer slice.Clear()

			slice.Sort(tt.algorithm)

			for index, val := range slice {

				if tt.want[index] != val {
					t.Errorf("Sort func = got %v, want %v", val, tt.want[index])
				}

			}
		})
	}
}

func TestSortStringer(t *testing.T) {

	add := []string{"C", "D", "F", "A", "H", "B", "G", "E"}
	want := []string{"A", "B", "C", "D", "E", "F", "G", "H"}

	tests := []struct {
		name      string
		add       slices.Slice[string]
		want      slices.Slice[string]
		algorithm slices.SortAlgorithm[string]
	}{
		{
			name:      "default sort",
			add:       add,
			want:      want,
			algorithm: nil,
		},
		{
			name:      "bubble sort",
			add:       add,
			want:      want,
			algorithm: slices.BubbleSort[string],
		},
		{
			name:      "insertion sort",
			add:       add,
			want:      want,
			algorithm: slices.InsertionSort[string],
		},
		{
			name:      "merge sort",
			add:       add,
			want:      want,
			algorithm: slices.MergeSort[string],
		},
		{
			name:      "quick sort",
			add:       add,
			want:      want,
			algorithm: slices.QuickSort[string],
		},
		{
			name: "any sort",
			add:  add,
			want: want,
			algorithm: slices.AnySort[string](func(a, b string) int {
				return cmp.Compare(a, b)
			}),
		},
	}

	slice := slices.MakeSlice[string](0, 0)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			slice.Append(tt.add...)
			defer slice.Clear()

			slice.Sort(tt.algorithm)

			for index, val := range slice {
				if tt.want[index] != val {
					t.Errorf("Sort func = got byte: %v val: %s, want %v", []byte(val), val, tt.want[index])
				}
			}
		})
	}
}
