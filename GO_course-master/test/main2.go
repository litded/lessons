package main

import "fmt"

func main() {
	var base = [5]int{1, 2, 3, 4, 5}
	slice := base[1:3]
	for i := 0; i < 1000; i++ {
		fmt.Printf("%T is type length is %v and capcity is %v\n", slice, len(slice), cap(slice))
		slice = append(slice, i)
	}

}
