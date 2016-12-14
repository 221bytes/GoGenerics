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

func getSize(a interface{}) uintptr {
	switch reflect.ValueOf(a).Index(0).Kind() {
	case reflect.String:
		return unsafe.Sizeof("")
	case reflect.Struct:
		return reflect.TypeOf(reflect.ValueOf(a).Index(0)).Size()
	default:
		panic("invalid type")
	}
}

func DoReverse(a interface{}) {

	start := reflect.ValueOf(a).Index(0).UnsafeAddr()
	size := getSize(a)
	length := reflect.ValueOf(a).Len()

	ReverseGeneric(unsafe.Pointer(start), size, length)
}
