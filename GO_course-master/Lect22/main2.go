package main

import "fmt"

type ClassicFunction func(a int, b int) int

func GeneralFunc(f ClassicFunction) {
	fmt.Println(f(20, 30))
}

func AnotherGeneralFunction(a func(a int, b int) int) {
	fmt.Println(a(20, 30))
}

func main() {
	f := func(a int, b int) int {
		return a*b - 2
	}
	GeneralFunc(f)
}
