package main

import "fmt"

func main() {
	b := 255
	a := &b

	fmt.Printf("%T and %v\n", a, a)
	var c *int

	if c == nil {
		fmt.Println("Zero value is:", c)
		c = &b
		fmt.Println("New value is:", c)
	}

	intPntr := new(int)
	fmt.Printf("Type %T val %v Type pntr value %T Ptr Value %v\n", intPntr, intPntr, *intPntr, *intPntr)
	*intPntr = 25
	fmt.Printf("Type %T val %v Type pntr value %T Ptr Value %v\n", intPntr, intPntr, *intPntr, *intPntr)

}
