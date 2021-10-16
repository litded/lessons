package main

import "fmt"

func main() {
	n := 0
	for n <= 10 {
		//like while loop
		fmt.Println(n)
		n++
	}

	for {
		//infinte loop
		//like 'while True'
		fmt.Println("Infinite!!!")
		n++
		if n > 15 {
			break
		}
	}
}
