package main

import "fmt"

func main() {
	var total int
	fmt.Scan(&total)

	nums := make([]int, total)
	for i := range nums {
		var val int
		fmt.Scan(&val)
		nums[i] = val
	}
	//fmt.Println(nums)

	var start, stop int
	fmt.Scan(&start)
	fmt.Scan(&stop)

	var sum int
	for i := start - 1; i < stop; i++ {
		sum += nums[i]
	}
	fmt.Println(sum)
}
