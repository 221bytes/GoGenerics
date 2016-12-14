package generics

import "testing"

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
