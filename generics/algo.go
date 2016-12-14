package generics

import (
	"reflect"
	"unsafe"
)

func ReverseGeneric(start unsafe.Pointer, size uintptr, lenght int) {
	last := uintptr(start) + size*uintptr(lenght)
	for i := lenght/2 - 1; i >= 0; i-- {
		opp := unsafe.Pointer(last - size - size*uintptr(i))
		b := unsafe.Pointer(uintptr(start) + size*uintptr(i))
		*(*uintptr)(b), *(*uintptr)(opp) = *(*uintptr)(opp), *(*uintptr)(b)
	}
}

func DoReverse(a interface{}) {

	start := reflect.ValueOf(a).Index(0).UnsafeAddr()
	size := reflect.ValueOf(a).Index(0).Type().Size()
	length := reflect.ValueOf(a).Len()

	ReverseGeneric(unsafe.Pointer(start), size, length)
}
