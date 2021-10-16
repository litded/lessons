package main

import "fmt"

type Worker interface {
	Work()
}

type Person struct {
	name string
	age  int
}

func (p Person) Work() {
	fmt.Println("Person worked!")
}

func FindType(i interface{}) {
	switch t := i.(type) {
	case string:
		fmt.Println("This is string and value : ", i.(string))
	case int:
		fmt.Println("This is int and value: ", i.(int))
	case float64:
		fmt.Println("This is float64 and value: ", i.(float64))
	case Worker:
		t.Work()
	default:
		fmt.Println("Unknown type!")
	}
}

func main() {
	FindType("Petr")
	FindType(22)
	FindType(123.5)
	FindType(true)

	p := Person{"Bob", 22}
	FindType(p)
}
