package main

import (
	"fmt"
	"time"
)

func Numbers(chnlNums chan int) {
	for i := 0; i < 50; i++ {
		chnlNums <- i
	}

	time.Sleep(10 * time.Second)
	chnlNums <- 9999
	close(chnlNums)
}

func main() {
	chnlNums := make(chan int)
	go Numbers(chnlNums)

	for val := range chnlNums {
		fmt.Println("Recieved from channel:", val)
	}
	// for {
	// 	if val, ok := <-chnlNums; !ok {
	// 		fmt.Println(val, ok)
	// 		break
	// 	} else {
	// 		fmt.Println("Recieved from channel:", val)
	// 	}

	// }
}
