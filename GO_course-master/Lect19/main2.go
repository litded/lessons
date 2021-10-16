package main

import (
	"fmt"
	"sync"
	"time"
)

//WaitGroups --- goroutines counter

func Worker(i int, wg *sync.WaitGroup) {
	fmt.Println("Goroutine #", i, "Started!")
	time.Sleep(2 * time.Second)
	fmt.Println("Goroutine #", i, "Done!")
	wg.Done() //counter -= 1
}

func main() {
	goNums := 10
	var wg sync.WaitGroup //counter = 0
	for i := 0; i < goNums; i++ {
		wg.Add(1)         //counter += 1
		go Worker(i, &wg) //Worker(param, &counter)
	}
	wg.Wait() // counter == 0 ?
	fmt.Println("All goroutines ended!!!")
}
