package slice

import (
	"reflect"
	"unsafe"
)

// CleanDuplicate removes duplicate items in slice.
func CleanDuplicate[T comparable](slice []T) []T {
	tmp := make(map[T]bool)
	var ret []T
	for _, v := range slice {
		if _, ok := tmp[v]; !ok {
			tmp[v] = true
			ret = append(ret, v)
		}
	}
	return ret
}

// ByteFromString converts a string to byte slice.
// TODO: From go1.20, we should use new functions in unsafe package instead of using reflect.StringHeader.
func ByteFromString(s string) []byte {
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
	var b []byte
	t := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	t.Data = stringHeader.Data
	t.Len = stringHeader.Len
	t.Cap = stringHeader.Len
	return b
}

// ByteToString converts a string to byte slice.
// TODO: From go1.20, we should use new functions in unsafe package instead of using reflect.SliceHeader.
func ByteToString(b []byte) string {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	var s string
	t := (*reflect.StringHeader)(unsafe.Pointer(&s))
	t.Data = sliceHeader.Data
	t.Len = sliceHeader.Len
	return s
}

// RemoveWhere removes elements in slice s (type is []T) where those elements
// run true in checkFunc.
// e.g.
// []int{1, 2, 3} = RemoveWhere([]int{1, 2, 3, 4, 5}, func(i int) bool { return s > 3 })
// removes all elements greater than 3.
func RemoveWhere[T any](s []T, checkFunc func(T) bool) []T {
	var ret []T
	for _, k := range s {
		if !checkFunc(k) {
			ret = append(ret, k)
		}
	}
	return ret
}
