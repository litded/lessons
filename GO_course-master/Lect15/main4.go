package main

import "fmt"

func SuperDescriber(i interface{}) {
	fmt.Printf("Type is %T and value %v\n", i, i)
}

func main() {
	str := "Hello world!"
	SuperDescriber(str)

	intValue := 22
	SuperDescriber(intValue)

	std := struct {
		name string
	}{
		name: "Bob",
	}
	SuperDescriber(std)
}
