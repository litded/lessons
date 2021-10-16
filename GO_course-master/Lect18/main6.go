package main

import "fmt"

// 123 -> [(1 * 1) + (2 * 2) + (3 * 3)] + [(1*1*1) +(2*2*2) + (3*3*3)]

func calcSquares(number int, sqrChannel chan int) {
	summ := 0
	for number != 0 {
		d := number % 10
		summ += d * d
		number /= 10
	}
	sqrChannel <- summ
}

func calcCubes(number int, cubChannel chan int) {
	summ := 0
	for number != 0 {
		d := number % 10
		summ += d * d * d
		number /= 10
	}
	cubChannel <- summ

}

func main() {
	number := 100505812
	sqrChannel := make(chan int)
	cubChannel := make(chan int)

	go calcSquares(number, sqrChannel) //[(1 * 1) + (2 * 2) + (3 * 3)]
	go calcCubes(number, cubChannel)   //[(1*1*1) +(2*2*2) + (3*3*3)]

	square, cube := <-sqrChannel, <-cubChannel
	fmt.Println("Result is:", square+cube)
}
