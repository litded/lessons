package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Mult(a, b int) int {
	return a * b
}

func main() {
	fmt.Println("Main part....")
	expr := Add(2, 3) + Sub(2, 4) + Mult(12, 6)
	fmt.Println(expr)
}
