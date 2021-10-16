package main

import "fmt"

type Student struct {
	firstName string
	lastName  string
	age       int
}

type Professor struct {
	//anonymous fields
	string
	int
}

func main() {
	stdPtr := &Student{"Jon", "Snow", 12}
	fmt.Println("Name:", (*stdPtr).firstName)
	fmt.Println("Lastname:", stdPtr.lastName)

	prf := Professor{
		string: "Dctr. Brown",
		int:    55,
	}
	prf2 := Professor{"Prof.  Beuer", 71}
	fmt.Println(prf, prf2)
	//individual fields of anon.field.struct
	fmt.Println(prf2.string)
}
