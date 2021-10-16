package math

import "testing"

// func TestConstruct(t *testing.T) {
// 	s := New(1, 2, 3)
// 	if s.isEmpty {
// 		t.Error("Set is empty!")
// 	}
// }

func BenchmarkFib10(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		Fib(10)
	}
}
