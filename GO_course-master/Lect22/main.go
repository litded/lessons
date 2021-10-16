package main

import (
	"fmt"
	"time"
)

func main() {
	go func(n string) {
		fmt.Println("Hello world!", n)
	}("Golang")
	time.Sleep(2 * time.Second)
	fmt.Println("Main finish")
}
