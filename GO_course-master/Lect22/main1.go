package main

import "fmt"

type Addition func(a int, b int) int

func main() {
	var a Addition = func(a, b int) int {
		return a + b - 1
	}
	res := a(10, 20)
	fmt.Println("Result is:", res)
	fmt.Printf("Type %T and Value %v\n", a, a)
}
