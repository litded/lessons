package main

import "fmt"

func PrintInfo(name *string, lastname *string) {
	defer fmt.Println("Defered Message!")
	if name == nil {
		panic("runtime error: name cannot be nil")
	}
	if lastname == nil {
		panic("runtime error: lastname cannot be nil")
	}
	fmt.Println(*name, *lastname)
	fmt.Println("PrintInfo() Done!")
}

func main() {
	name := "Bob"
	PrintInfo(&name, nil)
	fmt.Println("Function Main finished!")
}
