package main

import "fmt"

type Student struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	std1 := Student{
		lastName:  "Petrov",
		firstName: "Petr",
		age:       21,
	}

	std2 := Student{"Ivan", "Ivanov", 22}

	std3 := struct {
		firstName string
		lastName  string
	}{
		firstName: "Dima",
		lastName:  "Dmitriev",
	}

	fmt.Println("First Student:", std1)
	fmt.Println("Second Stduent: Name=", std2.firstName, "Lname:", std2.lastName, "Age:", std2.age)
	fmt.Println("Anon struct:", std3)
}
