package main

import "fmt"

func main() {
	var a [3]int
	a[0] = 20
	a[1] = 30
	a[2] = 40
	//b := [3]int{20, 30, 40}
	c := [...]int{12}
	fmt.Println("Arr ints is :", a)
	fmt.Println("Arr inst b is :", c)
	//d := [5]int{}

	city := [2]string{"Moscow", "SPB"}
	another := city
	another[0] = "Monaco"
	fmt.Println(city, another, "Length is:", len(another))

	for i := 0; i < len(another); i++ {
		fmt.Println("City is:", another[i])
	}

	for i, v := range another {
		fmt.Printf("%d index and value %v\n", i, v)
	}

	multDim := [3][2]string{{"a", "b"}, {"c", "d"}, {"e", "f"}}
	for _, v1 := range multDim {
		for _, v2 := range v1 {
			fmt.Printf("%v ", v2)
		}
		fmt.Printf("\n")
	}

}
