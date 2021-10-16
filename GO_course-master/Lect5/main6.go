package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var words [5]string
	input := bufio.NewScanner(os.Stdin)
	for i := 0; i < len(words); i++ {
		var message string
		input.Scan()
		message = input.Text() //strconv.ParseFloat(input.Text(), 32) // strconv.Atoi("12")

		words[i] = message
	}

	var inputs int
	input.Scan()
	inputs, _ = strconv.Atoi(input.Text())
	//for i:=0 ;i < inputs; i++{
	//.......
	//}

	fmt.Println("Messages:", words)
}
