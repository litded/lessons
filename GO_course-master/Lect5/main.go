package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("Initial array:", a)
	a[0] = 10
	a[1] = 20
	a[2] = 30
	a[3] = 40
	a[4] = 50
	//a[5] = 2000
	//len(a) - 1 == lastElemIndex
	fmt.Println("After assign:", a)
	//Short-hand array declaration
	b := [3]int{1, 2, 3}
	c := [4]int{200, 300}
	fmt.Println("Arr b:", b)
	fmt.Println("Arr c:", c)

	//Dot-size arr
	dotArr := [...]string{"Bob", "Alice", "Derek"}
	fmt.Printf("%T type and values: %v and length: %v\n", dotArr, dotArr, len(dotArr))

	anotherArr := [4]string{}
	fmt.Printf("%T type and values: %v and length: %v\n", anotherArr, anotherArr, len(anotherArr))
	//anotherArr = dotArr
}
