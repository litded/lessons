package main

import "fmt"

func main() {
	var a chan int
	if a == nil {
		fmt.Println("Channel is nil!", a)
		a = make(chan int)
		fmt.Printf("Type of a chan is %T value %v\n", a, a)
	}
}
