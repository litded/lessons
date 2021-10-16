package main

import "fmt"

func SlicerPanicer() {
	defer fmt.Println("First defer in slicer!")
	defer fmt.Println("Second defer in slicer!")
	nums := []int{1, 2}
	fmt.Println(nums[10])
}

func main() {
	defer fmt.Println("Deffered print in main function!")
	SlicerPanicer()
	fmt.Println("Main Function finished!")
}
