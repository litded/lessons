package main

import (
	"Lect12/str"
	"fmt"
)

func PrintEmployee(empl str.Employee) {
	fmt.Println("Name:", empl.FirstName)
	fmt.Println("Lastname:", empl.LastName)
	fmt.Println("Age:", empl.Age)
}

func main() {
	empl := str.Employee{
		FirstName: "Bob",
		LastName:  "Snow",
		Age:       27,
	}

	PrintEmployee(empl)
	//empl.salary = 20000
}
