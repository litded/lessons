package main

import "fmt"

type Triangle struct {
	a, b, c int
}

type Rectangle struct {
	a, b int
}

func (t Triangle) Perimeter() int {
	return t.a + t.b + t.c
}

func (r *Rectangle) Perimeter() int {
	//r.a = 200
	return 2 * (r.a + r.b)
}

func (r *Rectangle) PrintRectangle() {
	fmt.Println("A side:", (*r).a)
	fmt.Println("B side:", r.b)
}

func main() {
	r := Rectangle{2, 2}
	t := Triangle{2, 3, 4}
	fmt.Println("Per of rect:", r.Perimeter())
	fmt.Println("Per of triang:", t.Perimeter())
	r.PrintRectangle()
}
