package main

import "fmt"

func main() {
	var (
		age, salary = 35, 10000
		name        = "Bob"
	)

	width, height := 25.2, 350.12

	//width, height := 123.12, 123123.12312

	fmt.Println("Привет! Я", name, "Age is:", age)
	fmt.Println("Salary:", salary)
	fmt.Println(width, height)

	width, height1 := 123.12, 127937123.12
	fmt.Println(width, height1)

	var a int = 25
	var b = 43.12

	a, b = 12, 1234.2
	fmt.Println(a, b)

	c, d := 13, "hello"
	c, d1 := 123, "ROfl"
	fmt.Println(c, d, d1)

	c = 400

}
