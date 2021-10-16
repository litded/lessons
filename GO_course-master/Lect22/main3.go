package main

import "fmt"

func Sample() func(a int, b int) int {
	f := func(a int, b int) int {
		return a + b
	}
	return f
}

func main() {
	f := Sample()
	fmt.Println(f(10, 20))
}
