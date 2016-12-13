package generics

import "testing"

func BenchmarkGenericStruct(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		testGenericStruct()
	}
}

func BenchmarkStruct(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		testStruct()
	}
}

func BenchmarkGenericString(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		testGenericString()
	}
}

func BenchmarkString(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		testString()
	}
}
