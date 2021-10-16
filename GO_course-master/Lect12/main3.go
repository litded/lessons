package main

import "fmt"

type Address struct {
	city   string
	street string
}

type Person struct {
	name string
	age  int
	//Nested structure
	address Address
}

func main() {
	p := Person{
		name: "Bob",
		age:  25,
		address: Address{
			city:   "Moscow",
			street: "Ladozhskaya",
		},
	}

	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("City:", p.address.city)
	fmt.Println("Street:", p.address.street)
}
