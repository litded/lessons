package main

import "fmt"

type Rectangle struct {
	a, b int
}

func Area(r Rectangle) {
	fmt.Println("Area Function for Rectangle:", r.a*r.b)
}

func (r Rectangle) Area() {
	fmt.Println("Area Method for rectangle:", r.a*r.b)
}

func Perimeter(r *Rectangle) {
	fmt.Println("Perimeter Function for rectangle:", 2*((*r).a+r.b))
}

func (r *Rectangle) Perimeter() {
	fmt.Println("Perimeter Method for rectangle:", 2*(r.a+r.b))
}

func main() {
	r := Rectangle{2, 2}
	Area(r)
	r.Area()

	p := &r
	p.Area()
	//Area(p) --- function with value works only with value

	p.Perimeter()
	Perimeter(p)
	r.Perimeter()
	//Perimeter(r) --- function with pointer works only with pointer
}
