package main

import "fmt"

//pre-requisite for G task
const iterate = 5

func main() {
	var number int
	var currentMax int = -10000
	var currentMin int = 9999999
	//12 123 14 15 10 2999 2
	for i := 0; i < iterate; i++ {
		fmt.Scan(&number)
		if number > currentMax {
			currentMax = number
		}
		if number < currentMin {
			currentMin = number
		}
	}

	fmt.Println("Current Max is:", currentMax)
	fmt.Println("Current Min is:", currentMin)

	var realNumber float64 = 123.1234123432
	ans := fmt.Sprintf("%.1f", realNumber)

	fmt.Printf("%.1f \n", realNumber)
	fmt.Println("With sprintf:", ans)
}
