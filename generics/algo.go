package generics

import "unsafe"

type Test struct {
	s string
	a int
}

func ReverseGeneric(start unsafe.Pointer, size uintptr, lenght int) {
	last := uintptr(start) + size*uintptr(lenght)
	for i := lenght/2 - 1; i >= 0; i-- {
		opp := unsafe.Pointer(last - size - size*uintptr(i))
		b := unsafe.Pointer(uintptr(start) + size*uintptr(i))
		*(*uintptr)(b), *(*uintptr)(opp) = *(*uintptr)(opp), *(*uintptr)(b)
	}
}

func testGenericStruct() {
	a := []Test{{s: "0"}, {s: "1"}, {s: "2"}, {s: "3"}, {s: "4"}}
	start := unsafe.Pointer(&a[0])
	size := unsafe.Sizeof(a[0])
	lenght := len(a)
	ReverseGeneric(start, size, lenght)
}

func testStruct() {
	a := []Test{{s: "0"}, {s: "1"}, {s: "2"}, {s: "3"}, {s: "4"}}

	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
}

func testString() {
	a := []string{"0", "1", "2", "3", "4"}

	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
}

func testGenericString() {
	a := []string{"0", "1", "2", "3", "4"}
	start := unsafe.Pointer(&a[0])
	size := unsafe.Sizeof(a[0])
	lenght := len(a)
	ReverseGeneric(start, size, lenght)
}
