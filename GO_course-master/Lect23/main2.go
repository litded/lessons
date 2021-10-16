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
	if reflect.TypeOf(q).Kind() == reflect.Struct {
		v := reflect.ValueOf(q)
		fmt.Println("Number of fields:", v.NumField()) //узнать количество полей
		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("Field : %d type %T value %v\n", i, v.Field(i), v.Field(i))
		}
	}
}

func main() {
	std := Student{"Bob", 21} // insert into Student values(Bob, 21)
	//empl := Employee{"Alice", 903, "Moscow", 100500} //insert into Employee values(Alice, 903, Moscow, 100500)
	//fmt.Println(insertQuery(std))
	insertQuery(std)
}
