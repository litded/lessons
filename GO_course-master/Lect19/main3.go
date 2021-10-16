package main

import (
	"fmt"
	"time"
)

func FirstDataBase(out1 chan string) {
	time.Sleep(2 * time.Second)
	out1 <- "Answer from DB1"
}

func SecondDataBase(out2 chan string) {
	time.Sleep(1 * time.Second)
	out2 <- "Answer from DB2"
}

func main() {
	out1 := make(chan string)
	out2 := make(chan string)
	go FirstDataBase(out1)
	go SecondDataBase(out2)
	select {
	case val := <-out1:
		fmt.Println(val)
	case val := <-out2:
		fmt.Println(val)
	}
}
