package main

import (
	"fmt"
	"time"
)

func testFunction() {
	fmt.Println("Finish testFunction!")
}

func main() {
	defer testFunction()
	fmt.Println("Main function started!")
	time.Sleep(2 * time.Second)
	fmt.Println("Main function finished")
}
