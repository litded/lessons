package main

import (
	"fmt"
	"runtime/debug"
)

func RecoveryMainGoroutine() {
	if rec := recover(); rec != nil {
		fmt.Println("Recovered:", rec)
	}
}

func RecoveryTestGoroutine(done chan bool) {
	if rec := recover(); rec != nil {
		fmt.Println("Recovered another goroutine:", rec)
		done <- true
		//debug.Stack()
		debug.PrintStack()
	}
}

func Calculator(a, b int) {
	defer RecoveryMainGoroutine()
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	done := make(chan bool)
	go Division(a, b, done)
	<-done
}

func Division(a int, b int, done chan bool) {
	defer RecoveryTestGoroutine(done)
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
	done <- true
}

func main() {
	Calculator(10, 0)
	fmt.Println("Main function done!")
}
