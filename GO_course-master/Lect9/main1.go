package main

import "fmt"

func main() {
	map1 := make(map[string]string)
	map1["one"] = "1"
	map2 := map1
	map2["one"] = "eins"
	map2["two"] = "zwei"

	fmt.Println("Map1:", map1)
	fmt.Println("Map2:", map2)
	if map1 == nil {
	}
}
