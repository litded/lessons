package main

import (
	"fmt"
	"sort"
)

func main() {
	var name []int
	if name == nil {
		fmt.Println("NIL!")
	}
	lastnames := []int{20, 300, 1, 2, 3}
	name = append(name, lastnames...)
	fmt.Println(name)

	sort.Ints(lastnames)
	fmt.Println("Sorted:", lastnames)
}
