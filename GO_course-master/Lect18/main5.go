package main

import (
	"fmt"
	"time"
)

func myGoroutine(done chan bool) {
	fmt.Println("myGoroutine SLEEP!")
	time.Sleep(7 * time.Second)
	fmt.Println("myGoroutine awake!")
	done <- true
}

func main() {
	done := make(chan bool)
	fmt.Println("Main Goroutine want to call myGoroutine")
	go myGoroutine(done)
	<-done
	fmt.Println("Main Goroutine terminated!")
}
