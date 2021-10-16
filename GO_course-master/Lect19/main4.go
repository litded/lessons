package main

import (
	"fmt"
	"time"
)

func Worker(chnl chan string) {
	time.Sleep(2500 * time.Millisecond)
	chnl <- "Worker done!"
}

func main() {
	chnl := make(chan string)
	go Worker(chnl)
	for {
		time.Sleep(500 * time.Millisecond)
		select {
		case val := <-chnl:
			fmt.Println("Get value:", val)
			return
		default:
			fmt.Println("no value recieved")
		}
	}
}
