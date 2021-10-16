package main

import (
	"fmt"
	"time"
)

func myGoroutine() {
	fmt.Println("Hello from Goroutine!")
}

func main() {
	go myGoroutine()

	fmt.Println("Hello from main!")
	time.Sleep(1 * time.Second)
}
