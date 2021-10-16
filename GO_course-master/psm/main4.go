package main

import (
	"fmt"
	"psm/str"
	_ "psm/str"
)

type Employee struct {
	firstName string
	lastName  string
	age       int
}

type MoreInfo struct {
	city string
	Employee
}

func PrintInfo(mi MoreInfo) {
	fmt.Println("city:", mi.city)
	fmt.Println("Name:", mi.firstName)
	fmt.Println("Lastname:", mi.lastName)
}

func main() {

	e1 := Employee{"Bob", "Dilan", 45}
	e2 := Employee{
		firstName: "Derek",
		lastName:  "Jons",
		age:       25,
	}

	mi := MoreInfo{
		city:     "Moscow",
		Employee: e1,
	}
	fmt.Println(e1, e2)
	fmt.Println(mi, mi.city, mi.lastName)
	PrintInfo(mi)

	newEmp := str.Employeer{
		Name:     "Bob",
		Lastname: "Dilan",
	}
	fmt.Println(newEmp, newEmp.Lastname)
}
