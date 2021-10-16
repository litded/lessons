package main

import "fmt"

type Student struct {
	name string
	year int
}

func (s Student) printInfo() {
	fmt.Println("Name:", s.name)
	fmt.Println("Course:", s.year)
}

func main() {
	std := Student{
		name: "Petr",
		year: 3,
	}

	std.printInfo()
}
