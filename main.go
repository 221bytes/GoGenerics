package main

import (
	"unsafe"

	"fmt"

	"github.com/221bytes/generics/generics"
)

func main() {
	a := []string{"0", "1", "2", "3", "4"}
	fmt.Println(a)
	start := unsafe.Pointer(&a[0])
	size := unsafe.Sizeof(a[0])
	lenght := len(a)
	generics.ReverseGeneric(start, size, lenght)
	fmt.Println(a)
}
