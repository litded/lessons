package main

import "fmt"

func Reverse(message string) {
	for _, v := range []rune(message) {
		defer fmt.Printf("%c", v)
	}
}

func main() {
	message := "Hello world!"
	fmt.Println("Original message:", message)
	fmt.Printf("Reversed message: ")
	Reverse(message)
	fmt.Println()
}
