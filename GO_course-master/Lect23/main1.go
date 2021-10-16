package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	name string
	age  int
}

type Employee struct {
	name   string
	id     int
	city   string
	salary int
}

func insertQuery(q interface{}) {
	t := reflect.TypeOf(q)  //снятие типа
	k := t.Kind()           //какого вида (глубокий тип)
	v := reflect.ValueOf(q) //снятие значения
	fmt.Println("Type:", t)
	fmt.Println("Kind:", k)
	fmt.Println("Value:", v)
}

func main() {
	std := Student{"Bob", 21} // insert into Student values(Bob, 21)
	//empl := Employee{"Alice", 903, "Moscow", 100500} //insert into Employee values(Alice, 903, Moscow, 100500)
	//fmt.Println(insertQuery(std))
	insertQuery(std)
}
