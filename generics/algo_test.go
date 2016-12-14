package generics_test

import (
	"testing"
	"unsafe"

	"github.com/221bytes/generics/generics"
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

func BenchmarkGenericFastStruct(b *testing.B) {
	for n := 0; n < b.N; n++ {
		testGenericFastStruct()
	}
}

func BenchmarkGenericSlowStruct(b *testing.B) {
	for n := 0; n < b.N; n++ {
		testGenericSlowStruct()
	}
}

func BenchmarkStruct(b *testing.B) {
	for n := 0; n < b.N; n++ {
		testStruct()
	}
}

func BenchmarkGenericFastString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		testGenericFastString()
	}
}

func BenchmarkGenericSlowString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		testGenericSlowString()
	}
}

func BenchmarkString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		testString()
	}
}

func testGenericFastStruct() {
	a := generateArrayStruct()
	start := unsafe.Pointer(&a[0])
	size := unsafe.Sizeof(a[0])
	lenght := len(a)
	generics.ReverseGeneric(start, size, lenght)
}

func testGenericSlowStruct() {
	a := generateArrayStruct()
	generics.DoReverse(a)
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
	generics.ReverseGeneric(start, size, lenght)
}

func testGenericSlowString() {
	a := []string{"0", "1", "2", "3", "4"}
	generics.DoReverse(a)
}
