package main

import (
	"fmt"
	"smthng/math"
)

func main() {
	s := math.New(1, 2, 3)
	fmt.Println(s.IsEmpty())
	fmt.Println(s)

}
