package main

import (
	"fmt"
	"time"
)

func Writer(intChan chan int) {
	for i := 0; i < 10; i++ {
		intChan <- i
		fmt.Println("Wrote", i, "in channel")
	}
	close(intChan)
}

func main() {
	intChan := make(chan int, 3)
	go Writer(intChan)
	time.Sleep(2 * time.Second)
	for v := range intChan {
		fmt.Println("Read form channel", v)
		time.Sleep(1 * time.Second)
	}
}
