package rectangle

import "fmt"

//A is variable for ...
const A, b = 20, 30

func init() {
	fmt.Println("rectmetr.go init function")
	fmt.Println("A var is:", A, "b var is:", b)
}

//Area is function for ......
func Area(width, length float64) float64 {
	return width * length
}

func innerArea(width, length float64) float64 {
	return width*length + 1
}
