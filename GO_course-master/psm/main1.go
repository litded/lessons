package main

import "fmt"

func changeVal(val int) {
	val += 20
}

func changeValWithPtr(val *int) {
	*val += 200
}

func reper() *int {
	i := 220
	return &i
}

func main() {
	b := 200
	a := &b
	fmt.Println("address of b is", a)
	fmt.Println("value of b is", *a)
	*a += 200
	fmt.Println("New value of b:", *a, b)

	c, d := 20, 20
	changeVal(c)
	changeValWithPtr(&d)
	fmt.Println("New value of c is", c)
	fmt.Println("New value of d is:", d)

	v := reper()
	fmt.Println("Value of reper() is:", *v)
	*v += 200
	fmt.Println("After mods:", *v)
}
