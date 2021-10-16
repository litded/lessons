package main

import "Lect16/employee"

func main() {
	empl := employee.New("Bob", "Dylan", 10000000, 55)
	empl.ShowInfo()
}
