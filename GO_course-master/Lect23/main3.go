package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 11
	intX := reflect.ValueOf(x).Int()
	fmt.Printf("Type %T value %v\n", intX, intX)

	y := "Message"
	strY := reflect.ValueOf(y).String()
	fmt.Printf("Type %T value %v\n", strY, strY)
}
