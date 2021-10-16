package main

import "fmt"

func main() {
	var a = [3]string{"Moscow", "SPB", "KZN"}
	b := a
	b[0] = "Monaco"

	fmt.Println("a arr is :", a)
	fmt.Println("b arr is :", b)

	fmt.Println("Length of a:", len(a))

	//iterating
	for i := 0; i < len(a); i++ {
		fmt.Println("Value:", a[i])
	}

	for i, v := range a {
		fmt.Println("Index:", i, "Value:", v)
	}
}
