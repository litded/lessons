package main

import "fmt"

func main() {
	city := []string{"Moscow", "SPB", "KZN"}
	newCity := make([]string, len(city)-1)

	copy(newCity, city)
	newCity[0] = "Monaco"
	fmt.Println("New city:", newCity[:2])
	fmt.Println("Old city:", city)
}
