package main

import "fmt"

func calc(a, b int) int {
	var totalPrice = a * b
	return totalPrice
}

func geometryMetrics(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = 2 * (length + width)
	return
}

func main() {
	fmt.Println("Calc result:", calc(5, 10))
	area, per := geometryMetrics(22.5, 12.5)
	fmt.Println("Area is:", area, "Perimeter is:", per)

	names := make(map[string]string)
	names["bob"] = "after"
	names["alice"] = "before"
	names["joseph"] = "asdsdf"
	fmt.Println(names)
	for k, v := range names {
		fmt.Println(k, v)
	}
}
