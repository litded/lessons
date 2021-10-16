package employee

import "fmt"

type employee struct {
	Name     string
	LastName string
	Salary   int
	age      int
}

func New(name string, lastname string, salary int, age int) employee {
	e := employee{}
	e.Name = name
	e.LastName = lastname
	e.Salary = salary
	e.age = age

	return e
}

func (e employee) ShowInfo() {
	fmt.Println("Name:", e.Name)
	fmt.Println("LastName:", e.LastName)
	fmt.Println("Salary:", e.Salary)
	fmt.Println("Age:", e.age)
}
