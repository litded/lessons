package main

import "fmt"

const total = 100

func main() {
	var inputValue int64
	fmt.Scan(&inputValue)

	if inputValue < (total - 50) {
		fmt.Println("1")
	} else if inputValue >= (total-50) && inputValue < total {
		fmt.Println("2")
	} else if inputValue >= total && inputValue%2 != 0 {
		fmt.Println("3")
	} else {
		fmt.Println("4")
	}

	switch {
	case inputValue < (total - 50):
		fmt.Println("1")
	case inputValue >= (total-50) && inputValue < total:
		fmt.Println("2")
	case inputValue >= total && inputValue%2 != 0:
		fmt.Println("3")
	default:
		fmt.Println("4")
	}

}
