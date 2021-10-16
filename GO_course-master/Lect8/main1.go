package main

import "fmt"

//FindNum for ...
func FindNum(num int, nums ...int) {
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "was found with id:", i)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in", nums)
	}
}

//Modify for ...
func Modify(sl ...int) {
	sl[0] = 20000
	sl = append(sl, 30000)
	fmt.Println(sl)
}

func main() {
	FindNum(10, 20, 30, 40, 50, 10, 2)
	FindNum(10, 20)
	FindNum(10)
	slice := []int{20, 30, 40}
	FindNum(20, slice...)

	fmt.Println("Slice before modifications:", slice)
	Modify(slice...)
	fmt.Println("Slice after modifications:", slice)

}
