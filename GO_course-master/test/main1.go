package main

import "fmt"

func main() {
	var a = [5]int{1, 2, 3, 4, 5}
	b := a[1:4]
	c := []int{1, 2, 3, 4}
	fmt.Printf("%T and %v\n", b, b)
	fmt.Println("Slice:", c, "Length:", len(c), "Capacity:", cap(c))

	darr := [5]int{100, 200, 300, 400, 500}
	dslice := darr[1:4]
	for i := range dslice {
		dslice[i] += 10
	}
	fmt.Println("After modifications:", darr)

	examSlice := []int{100, 200, 300}
	anotherSlice := examSlice
	anotherSlice[1] = 10000
	fmt.Println("examSlice:", examSlice, "anotherSlice:", anotherSlice)
}
