package main

import "fmt"

func FinishFunction() {
	fmt.Println("Finish function!")
}

func FindMax(nums []int) {
	defer FinishFunction() //defer call
	fmt.Println("Started FindMax()")
	maxNums := nums[0]
	for _, v := range nums {
		if v > maxNums {
			maxNums = v
		}
	}
	fmt.Println("Max number is:", maxNums)
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 10, 4}
	FindMax(nums)
	fmt.Println("Finish Main Function!")
}
