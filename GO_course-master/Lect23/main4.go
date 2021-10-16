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

type Kek struct {
	kek string
}

func insertQuery(q interface{}) string {
	if reflect.TypeOf(q).Kind() == reflect.Struct {
		structName := reflect.TypeOf(q).Name()
		query := fmt.Sprintf("insert into %s values(", structName)
		values := reflect.ValueOf(q)
		for i := 0; i < values.NumField(); i++ {
			switch values.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					//insert into Student values(1
					query = fmt.Sprintf("%s%d", query, values.Field(i).Int())
				} else {
					//insert into Student values(1, 2
					query = fmt.Sprintf("%s, %d", query, values.Field(i).Int())
				}

			case reflect.String:
				if i == 0 {
					//insert into Student values("Bob"
					query = fmt.Sprintf("%s\"%s\"", query, values.Field(i).String())
				} else {
					//insert into Student values("Bob", "work"
					query = fmt.Sprintf("%s, \"%s\"", query, values.Field(i).String())
				}

			default:
				fmt.Println("unsupported type")
				return ""
			}
		}
		query = fmt.Sprintf("%s)", query)
		return query
	}
	fmt.Println("unsupported type")
	return ""
}

func main() {
	std := Student{"Bob", 21}                        // insert into Student values(Bob, 21)
	empl := Employee{"Alice", 903, "Moscow", 100500} //insert into Employee values(Alice, 903, Moscow, 100500)
	//fmt.Println(insertQuery(std))
	kek := Kek{"kek"}

	fmt.Println(insertQuery(std))
	fmt.Println(insertQuery(empl))
	fmt.Println(insertQuery(kek))
}
