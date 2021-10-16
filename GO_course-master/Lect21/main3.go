package main

import "fmt"

func RecoverPrintInfo() {
	if rec := recover(); rec != nil {
		fmt.Println("recovered from", rec)
	}
}

func PrintInfo(name *string, lastname *string) {
	defer RecoverPrintInfo()
	if name == nil {
		panic("runtime error: name cannot be nil")
	}
	if lastname == nil {
		panic("runtime error: lastname cannot be nil")
	}
	fmt.Println(*name, *lastname)
	fmt.Println("PrintInfo() Done!")
}

func main() {
	name := "Bob"
	PrintInfo(&name, nil)
	fmt.Println("Function Main finished!")
}
