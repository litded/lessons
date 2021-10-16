package main

import "fmt"

type Student struct {
	name string
	age  int
}

func (s Student) changeName(newName string) {
	s.name = newName
}

func (s *Student) changeAge(newAge int) {
	s.age = newAge
}

func main() {
	std1 := Student{"Bob", 21}
	fmt.Println("New Student:", std1)
	std1.changeName("Bobby")
	std1.changeAge(22)
	fmt.Println("Stduent after changing methods:", std1)
}
