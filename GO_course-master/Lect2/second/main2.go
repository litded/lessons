package main

import (
	"fmt"
	"strings"
)

const total string = "чат"

//K solution

func main() {

	var message string

	fmt.Scan(&message)

	if strings.Contains(message, total) {
		fmt.Println("БОТ")
	} else {
		fmt.Println("НЕ БОТ")
	}
}
