package main

import "fmt"

type Address struct {
	city string
}

func (a Address) Sample() {
	fmt.Println("Sample method from Address")
}

type GlobalAddress struct {
	street string
	Address
}

func (ga GlobalAddress) Sample() {
	fmt.Println("Sample method from GlobalAddress")
}

func main() {
	ga := GlobalAddress{}
	ga.Sample()
	ga.Address.Sample()
}
