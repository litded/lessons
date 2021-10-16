package main

import "fmt"

//Named structure
type Employee struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	empl1 := Employee{
		firstName: "Bob",
		age:       25,
		lastName:  "Dub",
	}
	empl2 := Employee{"Alice", "Yandex", 11}

	//Anonymous structure
	empl3 := struct {
		firstName string
		lastName  string
	}{
		firstName: "Georg",
		lastName:  "Tomp",
	}

	fmt.Println(empl1, empl2, empl3)
	fmt.Println(empl1.age)

	var empl4 Employee
	fmt.Println(empl4.age)
}
