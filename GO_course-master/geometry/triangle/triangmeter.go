package triangle

import (
	"fmt"
	"math"
)

func init() {
	fmt.Println("triangmeter.go init started")
	//fmt.Println("Perimeter:", rectangle.Perimeter(20, 30))
}

//Perimeter is ....
func Perimeter(a, b, c float64) float64 {
	return a + b + c
}

//Area is .....
func Area(a, b, c float64) float64 {
	p := Perimeter(a, b, c) / 2
	result := math.Sqrt(p * (p - a) * (p - b) * (p - c))
	return result
}
