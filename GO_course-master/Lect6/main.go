package main

import "fmt"

func main() {
	arr := [5]int{10, 20, 30, 40, 50}
	var bSlice []int = arr[1:4]
	fmt.Printf("%T and value %v\n", bSlice, bSlice)
	c := []int{1, 2, 3}
	fmt.Println("Slice:", c)

	///Slice modifications
	fmt.Println("Arr before modifications:", arr)
	for i := range bSlice {
		bSlice[i] += 200
	}
	fmt.Println("Arr after modifications:", arr)

	//Slice is referal type
	cSlice := bSlice
	cSlice[0] = 10000000
	fmt.Println("cSlice:", cSlice)
	fmt.Println("bSlice:", bSlice)
	fmt.Println("Arr:", arr)

	// var d = []int{}
	// copy(d, bSlice)
	// fmt.Println("d slice:", d)
}
