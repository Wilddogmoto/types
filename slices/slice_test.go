package slices_test

import (
	"fmt"
	"github.com/Wilddogmoto/types/slices"
	"reflect"
	"sync"
	"testing"
)

func ExampleMakeSlice() {

	slice := slices.MakeSlice[int](0, 3)
	slice.Append(1, 2, 3)

	fmt.Print(slice)
	// Output: [1 2 3]
}

func TestSlice_Remove(t *testing.T) {

	index := 1

	slice := slices.MakeSlice[int](2, 2)

	slice.AppendIndex(100, index)
	slice.Remove(index)

	val := slice[0]
	if val != 0 {
		t.Error("error Remove func")
	}
}

func TestSlice_Copy(t *testing.T) {

	slice := slices.MakeSlice[int](0, 0)
	slice.Append(1, 2, 3)

	newslice := slice.Copy()

	if reflect.ValueOf(newslice).Pointer() == reflect.ValueOf(slice).Pointer() {
		t.Error("error Copy func")
	}
}

func TestSlice_Clear(t *testing.T) {

	slice := slices.MakeSlice[int](0, 5)
	slice.Append(1, 2, 3, 4, 5)

	slice.Clear()

	if slice.Len() != 0 || slice.Cap() != 5 {
		t.Error("error Clear func")
	}

}

func TestSlice_Grow(t *testing.T) {

	slice := slices.MakeSlice[int](0, 5)
	slice.Append(1, 2, 3, 4, 5)

	slice.Grow(5)

	if slice.Cap() != 10 {
		t.Error("error Grow func")
	}

}

func TestSearchInteger(t *testing.T) {

	want := 100

	add := []int{1, 3, want, 33, 500, 10, 0}

	tests := []struct {
		name string
		add  slices.Slice[int]
		want int
	}{
		{
			name: "default sort",
			add:  add,
			want: want,
		},
	}

	slice := slices.MakeSlice[int](0, 10)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			slice.Append(tt.add...)
			slice.Sort(nil)
			defer slice.Clear()

			index, find := slice.Search(tt.want)

			if !find {
				t.Error("Search func value not found")
				return
			}

			if slice[index] != tt.want {
				t.Errorf("Search func = got %v, want %v", slice[index], tt.want)
			}
		})
	}
}

func TestSearchStringer(t *testing.T) {

	want := "find"

	add := []string{"Alice", "Bob", want, "Some_name"}

	tests := []struct {
		name string
		add  slices.Slice[string]
		want string
	}{
		{
			name: "default sort",
			add:  add,
			want: want,
		},
	}

	slice := slices.MakeSlice[string](0, 10)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			slice.Append(tt.add...)
			slice.Sort(nil)
			defer slice.Clear()

			index, find := slice.Search(tt.want)

			if !find {
				t.Error("Search func value not found")
				return
			}

			if slice[index] != tt.want {
				t.Errorf("Search func = got %v, want %v", slice[index], tt.want)
			}
		})
	}
}

func TestAsyncSlice(t *testing.T) {

	capacity := 100

	slice := slices.MakeSlice[int](capacity, capacity).Async()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {

		defer wg.Done()
		for i := 0; i != capacity; i++ {
			t.Logf("iter first go: %d", i)
			if slice.Len() == i {
				return
			}
			slice.Append(i)
		}
	}()

	wg.Add(1)
	go func() {

		defer wg.Done()
		for i := 0; i != capacity; i++ {
			t.Logf("iter two go: %d", i)
			if slice.Len() == i {
				return
			}

			slice.Append(i)
		}
	}()

	wg.Wait()
}
