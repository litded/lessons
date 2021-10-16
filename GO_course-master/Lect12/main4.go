package main

import "fmt"

type Address struct {
	city   string
	street string
}

type Person struct {
	name string
	age  int
	//Promoted structure
	Address
}

func main() {
	p := Person{
		name: "Bob",
		age:  25,
		Address: Address{
			city:   "Moscow",
			street: "Ladozhskaya",
		},
	}

	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("City:", p.city)
	fmt.Println("Street:", p.street)
}
