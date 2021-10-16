package main

// L solution

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	var login, email string
	fmt.Scan(&login)
	fmt.Scan(&email)

	if utf8.RuneCountInString(login) < 10 || strings.Contains(login, "@") {
		fmt.Println("Некорректный логин")
	} else if !(strings.Contains(email, "@") && strings.Contains(email, ".")) {
		fmt.Println("Некорректная почта")
	} else {
		fmt.Println("ОК")
	}
}
