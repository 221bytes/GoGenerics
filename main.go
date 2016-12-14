package main

import (
	"fmt"
	"reflect"
	"unsafe"

	"github.com/221bytes/generics/generics"
)

type Test struct {
	s string
	a int
}

func main() {
	a := []Test{{s: "0"}, {s: "1"}, {s: "2"}, {s: "3"}, {s: "4"}, {s: "5"}}
	// a := []string{"0", "1", "2", "3", "4", "5"}

	fmt.Println(a)
	normalStart := unsafe.Pointer(&a[0])
	normalSize := unsafe.Sizeof(a[0])
	normalLenght := len(a)

	start := reflect.ValueOf(a).Index(0).UnsafeAddr()
	size := reflect.TypeOf(reflect.ValueOf(a).Index(0)).Size()
	lenght := reflect.ValueOf(a).Len()

	fmt.Printf("%v == %v\n", normalStart, unsafe.Pointer(start))
	fmt.Printf("%v == %v\n", normalSize, size)
	fmt.Printf("%v == %v\n", normalLenght, lenght)
	generics.DoReverse(a)
	fmt.Println(a)

	generics.ReverseGeneric(normalStart, normalSize, normalLenght)
	fmt.Println(a)
}
