package main

import "fmt"

const (
	aConst = 25
	bConst = 27
)

func main() {
	const cConst = 75

	age := 26
	var a int
	if fmt.Scan(&a); age*a > 10000 {
		fmt.Println("BOOMER")
	} else {
		fmt.Println("ZOOMER")
	}

}
