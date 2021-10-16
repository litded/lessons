package main

import "fmt"

func main() {
	//employee := make(map[string]int)
	//var employee map[string]int
	employee := map[string]int{
		"Steve": 43,
		"Alex":  66,
	}
	employee["Bob"] = 25
	employee["Alice"] = 38
	employee["Jo"] = 18
	employee["Georg"] = 0
	employee["Petya"] = 33

	fmt.Println("Default map:", employee)
	fmt.Println("Alice age:", employee["Alice"])

	fmt.Println("Map length:", len(employee))

	fmt.Println("Georg is :", employee["Georg"], employee)

	if value, ok := employee["Petya"]; ok {
		fmt.Println(value)
	}

	delete(employee, "Petya")

	for key, val := range employee {
		fmt.Println("Key:", key, "Value:", val)
	}

}
