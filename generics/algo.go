package generics

import (
	"reflect"
	"unsafe"
)

type Test struct {
	s string
	a int
}

var a []Test
var s []string

func init() {
	a = generateArrayStruct()
	s = generateArrayString()
}

func generateArrayStruct() []Test {
	a := make([]Test, 1000000)
	for i := 0; i < 1000000; i++ {
		a[i] = Test{s: string(i)}
	}
	return a
}

func generateArrayString() []string {
	a := make([]string, 1000000)
	for i := 0; i < 1000000; i++ {
		a[i] = string(i)
	}
	return a
}

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

func testGenericFastStruct() {
	a := generateArrayStruct()
	start := unsafe.Pointer(&a[0])
	size := unsafe.Sizeof(a[0])
	lenght := len(a)
	ReverseGeneric(start, size, lenght)
}

func testGenericSlowStruct() {
	a := generateArrayStruct()
	DoReverse(a)
}

func testStruct() {
	a := generateArrayStruct()
	length := len(a)
	for i := length/2 - 1; i >= 0; i-- {
		opp := length - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
}

func testString() {
	a := []string{"0", "1", "2", "3", "4"}
	length := len(a)
	for i := length/2 - 1; i >= 0; i-- {
		opp := length - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
}

func testGenericFastString() {
	a := []string{"0", "1", "2", "3", "4"}
	start := unsafe.Pointer(&a[0])
	size := unsafe.Sizeof(a[0])
	lenght := len(a)
	ReverseGeneric(start, size, lenght)
}

func testGenericSlowString() {
	a := []string{"0", "1", "2", "3", "4"}
	DoReverse(a)
}
