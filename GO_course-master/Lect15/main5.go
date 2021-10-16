package main

import "fmt"

func GuessType(i interface{}) {
	gValue, ok := i.(int)
	if !ok {
		fmt.Println("Value of this type not found!")
	}
	fmt.Println(gValue)
}

func main() {
	var a interface{} = 12.12
	GuessType(a)

	var b interface{} = 12
	GuessType(b)
}
