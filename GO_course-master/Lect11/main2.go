package main

import "fmt"

func arrChanger(arr *[3]int) {
	//(*arr)[0] = 500
	arr[0] = 500
}

func sliceChanger(sls []int) {
	sls[0] = 5000
}

func main() {
	nums := [3]int{1, 2, 3}
	fmt.Println("Old nums:", nums)

	//bad idea
	arrChanger(&nums)
	fmt.Println("New nums:", nums)

	//go way ideomatic
	sliceChanger(nums[:])
	fmt.Println("After sliceChanger:", nums)
}
