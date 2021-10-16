package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := -n; i <= n; i++ {
		fmt.Println("Value:", i, "Squared:", i*i)
	}
}
