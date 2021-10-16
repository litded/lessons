package main

import "fmt"

func Add(a int) func(b int) int {
	Add2 := func(b int) int {
		return b + a
	}
	return Add2
}

func main() {
	f := Add(2) // func(b int) int {
	//	return b + 2
	//}
	f(3)
	fmt.Println(Add(2)(3))
}
