package main

import "fmt"

type MyInt int

func (a MyInt) Power(b int) MyInt {
	var result MyInt = 1
	for i := 0; i < b; i++ {
		result *= a
	}
	return result
}

func main() {
	var a MyInt = MyInt(2)
	fmt.Println(a.Power(4))
}
