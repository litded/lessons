package main

import (
	"fmt"
	_ "geometry/circle"
	"geometry/rectangle"
	"geometry/triangle"
)

func init() {
	fmt.Println("geometry.go init function")
}

func main() {

	fmt.Println("Geometry.go/mainFunc Started here")
	w, l := 22.5, 43.2
	area := rectangle.Area(w, l)
	perimeter := rectangle.Perimeter(w, l)

	fmt.Println("Area from rectangle:", area)
	fmt.Println("Perimeter from rectangle:", perimeter)
	//fmt.Println("innerArea(20,30):", rectangle.innerArea(20, 30))
	fmt.Println("A exproted:", rectangle.A)
	//fmt.Println("b unexported:", rectangle.b)

	a, b, c := 12.5, 13.5, 25.5
	fmt.Println("Area of triangle:", triangle.Area(a, b, c))
	fmt.Println("Perimeter of triangle:", triangle.Perimeter(a, b, c))
}
