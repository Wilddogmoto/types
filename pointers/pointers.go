package pointers

import (
	"cmp"
	"unsafe"
)

func GetOrderedPointer[T cmp.Ordered](value T) *T {
	return &value
}

func GetPointer[T any](value T) *T {
	return &value
}

func GetAddressPointer[T any](value *T) uintptr {
	return uintptr(unsafe.Pointer(value))
}

func GetValueByAddress[T any](ptr uintptr) T {
	return *(*T)(unsafe.Pointer(ptr))
}
