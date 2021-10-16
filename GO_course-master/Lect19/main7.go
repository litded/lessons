package main

import (
	"fmt"
	"sync"
)

var num = 0

func SuperIncrementer(wg *sync.WaitGroup, mut *sync.Mutex) {
	mut.Lock()
	num = num + 2
	mut.Unlock()
	wg.Done()
}

func main() {
	var mut sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go SuperIncrementer(&wg, &mut)
	}
	wg.Wait()
	fmt.Println("Result is:", num)
}
