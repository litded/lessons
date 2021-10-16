package main

import (
	"fmt"
	"sort"
)

func main() {
	var a = []int{300, 400, 500, -1, 20, 40}
	sort.Ints(a)
	fmt.Println(a)
}
