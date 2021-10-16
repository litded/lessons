package main

import "fmt"

type Student struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	var std Student
	var std2 = Student{
		age: 54,
	}
	fmt.Println(std)
	std.firstName = "Bob"
	std.lastName = "Jack"
	std.age = 12

	fmt.Println(std)
	fmt.Println(std2)
}
