package main

import "fmt"

func main() {
	var name = "Gello"
	runes := []rune(name)
	runes[0] = 'H'
	//name = string(runes)
	fmt.Println(string(runes), name)
}
