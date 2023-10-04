package pointers

import (
	"cmp"
)

func GetOrderedPointer[T cmp.Ordered](value T) *T {
	return &value
}

func GetPointer[T any](value T) *T {
	return &value
}
