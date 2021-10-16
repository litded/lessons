package main

import "fmt"

//variadic parameter a ...int
func calcSum(a ...int) int {
	fmt.Printf("%T type and %v values\n", a, a)
	return 2
}

func main() {
	slc := []string{"1", "2", "3"}
	slc = append(slc, "word")
	slc = append(slc, "word1", "word2")
	slc = append(slc, "word4", "word5", "word6")
	fmt.Println(slc, slc)
	_ = calcSum(20, 30, 40, 50)

}
