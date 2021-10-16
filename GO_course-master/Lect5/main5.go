package main

import "fmt"

//solution H

func main() {
	var sum int // initial sum
	var count int
	var nums [20]int //[0,0,0,0,0,....]

	fmt.Scan(&count)
	for i := 0; i < count; i++ {
		var value int
		fmt.Scan(&value)
		nums[i] = value
	}
	var start, stop int
	fmt.Scan(&start)
	fmt.Scan(&stop)

	for i := start - 1; i < stop; i++ {
		sum += nums[i]
	}
	fmt.Println(sum)
}
