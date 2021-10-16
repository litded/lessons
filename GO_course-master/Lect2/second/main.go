package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var a, b float32 = 22.5, 23.5

	//CLASSIC IF-ELSE
	if int(a+b)%2 == 0 {
		fmt.Println("ЧЁТНОЕ")
	} else {
		fmt.Println("НЕЧЁТНОЕ")
	}

	//CLASSIC IF-ELSE IF -- ELSE
	var age = 27
	if age < 14 {
		fmt.Println("1")
	} else if age >= 14 && age < 21 {
		fmt.Println("2")
	} else if age >= 21 && age < 33 {
		fmt.Println("3")
	} else {
		fmt.Println("4")
	}

	utf8.RuneCountInString(message) //количество символов в строке message

	//!(true && true) == !true || !true

}
