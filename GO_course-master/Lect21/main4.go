package main

import (
	"fmt"
	"runtime/debug"
)

func RecoverSlicerPanic() {
	if rec := recover(); rec != nil {
		fmt.Println("Recovered:", rec)
		debug.PrintStack()
	}
}

func SlicerPanicer() {
	defer fmt.Println("First defer in slicer!")
	defer fmt.Println("Second defer in slicer!")
	defer RecoverSlicerPanic()
	nums := []int{1, 2}
	fmt.Println(nums[10])
}

func main() {
	defer fmt.Println("Deffered print in main function!")
	SlicerPanicer()
	fmt.Println("Main Function finished!")
}
