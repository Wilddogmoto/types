package pointers

import (
	"cmp"
)

func GetPointer[T cmp.Ordered](value T) *T {
	return &value
}
