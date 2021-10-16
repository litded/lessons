package main

import (
	"fmt"
	"time"
)

func Numbers() {
	for i := 0; i <= 10; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func Letters() {
	for i := 'a'; i <= 'j'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

func main() {
	go Numbers()
	go Letters()
	time.Sleep(5 * time.Second)
	fmt.Println("Main Goroutine termonated :(")
}
