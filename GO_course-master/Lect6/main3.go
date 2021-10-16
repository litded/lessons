package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	newNumbers := []int{10, 20, 30, 40}

	langs := [][]string{
		{"C++", "C", "RUST", "Lua"},
		{"JS", "Python", "PHP"},
	}
	numbers = append(numbers, newNumbers...)
	// for _, v := range newNumbers {
	// 	numbers = append(numbers, v)
	// }
	fmt.Println("Numbers modification:", numbers)
}
