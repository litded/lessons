package main

import (
	"fmt"
)

func myGoroutine(done chan bool) {
	fmt.Println("Hello from Goroutine!")
	done <- true
}

func main() {
	done := make(chan bool)
	go myGoroutine(done)
	<-done
	fmt.Println("Hello from main!")

}
