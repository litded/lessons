package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	fmt.Println("Message is:", input.Text())
	kek := input.Text()
	fmt.Printf("%T and %v length is %v and cap is %v\n", kek, kek, len(kek), cap([]byte(kek)))
}
