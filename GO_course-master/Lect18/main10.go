package main

import "fmt"

// 123 -> [(1 * 1) + (2 * 2) + (3 * 3)] + [(1*1*1) +(2*2*2) + (3*3*3)]
func Nums(number int, numsChannel chan int) {
	for number != 0 {
		d := number % 10
		numsChannel <- d
		number /= 10
	}
	close(numsChannel)
}

func calcSquares(number int, sqrChannel chan int) {
	summ := 0
	nChan := make(chan int)
	go Nums(number, nChan)
	for d := range nChan {
		summ += d * d
	}
	sqrChannel <- summ
}

func calcCubes(number int, cubChannel chan int) {
	summ := 0
	nChan := make(chan int)
	go Nums(number, nChan)
	for d := range nChan {
		summ += d * d * d
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
