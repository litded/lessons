package main

import "fmt"

func main() {
	arr := [5]int{2, 3, 4, 5, 6}
	slice := arr[1:3]
	slice = append(slice, 10, 11, 12)

	for i := 0; i < 1000; i++ {
		slice = append(slice, i)
		fmt.Printf("%T and length %v and capacity %v\n", slice, len(slice), cap(slice))
	}

}
