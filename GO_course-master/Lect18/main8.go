package main

import "fmt"

func Sender(sendChannel chan<- int) {
	sendChannel <- 200
}

func main() {
	sendChannel := make(chan int)
	go Sender(sendChannel)
	fmt.Println(<-sendChannel)
}
