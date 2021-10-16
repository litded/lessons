package main

import "fmt"

type InnerFunction func(b int) int

func Adder(a int) InnerFunction {
	var inner InnerFunction = func(b int) int {
		return b + a
	}
	return inner
}

func Adder10(b int) int {
	return 10 + b
}
func Adder11(b int) int {
	return 11 + b
}

func main() {
	a := Adder(11)
	fmt.Println(Adder(11)(5))
	for i := 0; i < 10; i++ {
		fmt.Println(a(i), Adder11(i))
	}

	workers := []InnerFunction{a, a, a}
}
