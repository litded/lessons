package main

import "fmt"

func main() {
	var sum int
	var nums [5]int
	fmt.Println("Length:", len(nums), "Capacity:", cap(nums))

	for i := 0; i < 5; i++ {
		var temp int
		fmt.Scan(&temp)
		nums[i] = temp
		sum += temp
	}

	fmt.Printf("Arr:%v Type:%T Len:%v\n", nums, nums, len(nums))
	fmt.Println("Sum of all elements:", sum)

	fmt.Println("Length:", len(nums), "Capacity:", cap(nums))

}
