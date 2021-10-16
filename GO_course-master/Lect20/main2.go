package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func (s Student) PrintInfo() {
	fmt.Println("Student name is:", s.Name, "Student age is:", s.Age)
}

func main() {
	std := Student{"Bob", 21}
	defer std.PrintInfo()
	fmt.Println("Main Function Finished")
}
