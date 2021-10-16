package main

import (
	"fmt"
	"testing"
)

func TestMult(t *testing.T) {
	fmt.Println("Test Mult started")
	expect := 10
	res := Mult(2, 5)
	if res != expect {
		t.Error("Failed")
	}
}

func TestSomething(t *testing.T) {
	fmt.Println("Test something else")
	fmt.Println("This shouldnt run with -run=Add")
}
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(2, 5)
	}
}

func BenchmarkSub(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sub(10, 2)
	}
}

func benchmarkMult(a int, c int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		Mult(a, c)
	}
}

func BenchmarkMult100(b *testing.B) { benchmarkMult(100, 100, b) }
func BenchmarkMult25(b *testing.B)  { benchmarkMult(25, 25, b) }
