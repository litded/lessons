package main

import "fmt"

func PrintVal(val int) {
	fmt.Println("Value is:", val)
}

func main() {
	val := 10
	defer PrintVal(val)
	defer PrintVal(val + 1)
	defer PrintVal(val + 2)
	val = 20
	fmt.Println("Main function done with val:", val)
}
