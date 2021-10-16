package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	message := "Привет мир!"
	fmt.Printf("Bytes: ")
	for _, letter := range message {
		fmt.Printf("%x ", letter)
	}
	fmt.Printf("\nChars: ")

	runes := []rune(message) //from string to slice rune

	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
	fmt.Printf("\n")

	byteSlice := []byte{66, 67, 102, 195}
	str := string(byteSlice)
	fmt.Println("From byte slice:", str)

	fmt.Printf("String Bob len %v and runes %v\n", len("Bob"), utf8.RuneCountInString("Bob"))
	fmt.Printf("String Коля len %v and runes %v\n", len("Коля"), utf8.RuneCountInString("Коля"))
}
