package main

import "fmt"

type Circle struct {
	r int
}

type Rectangle struct {
	a, b int
}

type Triangle struct {
	a, b, c int
}

func (c *Circle) Perimeter() float64 {
	return 3.14 * float64(c.r) * 2
}

func (r *Rectangle) Perimeter() float64 {
	return float64(2 * (r.a + r.b))
}

func (r *Rectangle) Area() float64 {
	return float64(r.a * r.b)
}

func (t *Triangle) Perimeter() float64 {
	return float64(t.a + t.b + t.c)
}

type Figure interface {
	Perimeter() float64
}

func SummaryPerimeter(figures []Figure) float64 {
	var total float64
	for _, fig := range figures {
		total += fig.Perimeter()
	}
	return total
}

func main() {

	rect := Rectangle{1, 2}
	circle := Circle{5}
	triang := Triangle{3, 4, 5}

	figures := []Figure{&rect, &circle, &triang}

	fmt.Println("Result:", SummaryPerimeter(figures))
}
