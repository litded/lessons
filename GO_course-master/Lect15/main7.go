package main

import "fmt"

type Duck interface {
	Quack()
}

type Dog interface {
	GoodBoy()
}

type CCOMBO interface {
	Duck
	Dog
}

type MainChar struct {
	id   int
	exp  int
	name string
}

func (mc MainChar) Quack() {
	fmt.Println("Main Char QUACKED!")
}

func (mc MainChar) GoodBoy() {
	fmt.Println("Main Char is a Good boy!")
}

func main() {
	mc := MainChar{1, 200, "ROBO7000"}
	var dog Dog = mc
	dog.GoodBoy()

	var duck Duck = mc
	duck.Quack()

	var ccombo CCOMBO = mc
	ccombo.GoodBoy()
	ccombo.Quack()
}
