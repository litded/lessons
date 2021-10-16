package main

import "fmt"

type Duck interface {
	Quack()
}

type RoboDuck struct {
	name string
	id   int
}

func (rd RoboDuck) Quack() {
	fmt.Println("ROBO-QUACK! from:", rd.name)
}

//interface (type, value)

func Describer(d Duck) {
	fmt.Printf("Interface type is %T and value %v\n", d, d)
}

type PnevmoDuck struct{}

func (pd PnevmoDuck) Quack() {
	fmt.Println("PNEVMO-QUCK!")
}

func main() {
	rd := RoboDuck{"PERCEP9000", 12}
	pd := PnevmoDuck{}
	var duck Duck = rd
	var duck2 Duck = pd

	var duck3 Duck
	Describer(duck3)

	if duck3 == nil {
		fmt.Println("NIIIL!")
	}
	//duck3.Quack()

	Describer(duck2)
	Describer(duck)
	duck.Quack()
}
