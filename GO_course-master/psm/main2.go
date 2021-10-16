package main

import "fmt"

func mods(a *[3]int) {
	a[0] = 90
}

func modsSlice(a []int) {
	a[0] = 900
}

func main() {
	a := [3]int{1, 2, 3}
	fmt.Println(a)
	mods(&a)
	fmt.Println(a)
	modsSlice(a[:])
	fmt.Println(a)
}
