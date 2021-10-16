package main

import "fmt"

func main() {
	slice := make([]int, 0, 6)
	if slice == nil {
		fmt.Println("Default also nil!")
	}
	fmt.Println(slice)

	var names []string //default value for []string is nil!
	if names == nil {
		fmt.Println("Default is nil!")
	}
	fmt.Printf("%v\n", names)
	names = append(names, "Bob")
	fmt.Println(names)
}
