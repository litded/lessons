package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
	"unsafe"
)

func main() {
	//Logic Type
	fmt.Println("===============BOOLEAN===================")
	var t, f bool = true, false
	fmt.Println("AND:", t && f)
	fmt.Println("OR:", t || f)
	fmt.Println("NOT:", !t)
	fmt.Printf("%T and %v and %v bytes\n", t, t, unsafe.Sizeof(t))
	fmt.Println()

	fmt.Println("======================INTEGERS=====================")
	var newInt int = 12
	var (
		aInt8  int8  = 125
		aInt16 int16 = 200
		aInt64 int64 = 12039
	)

	fmt.Printf("%T nad %v and %v bytes for newInt\n", newInt, newInt, unsafe.Sizeof(newInt))
	fmt.Printf("%T nad %v and %v bytes for aInt8\n", aInt8, aInt8, unsafe.Sizeof(aInt8))
	fmt.Printf("%T nad %v and %v bytes for aInt16\n", aInt16, aInt16, unsafe.Sizeof(aInt16))
	fmt.Printf("%T nad %v and %v bytes for aInt64\n", aInt64, aInt64, unsafe.Sizeof(aInt64))
	/*
		var bInt8, cInt8 int8 = 127, 127
		dInt8 := bInt8 + cInt8
		fmt.Println(dInt8)
	*/
	//Line commet
	var age int = 25
	var salary int64 = 200000
	fmt.Println("Solution:", int64(age)*salary)

	var lhs, rhs int = 22, 33
	fmt.Println("Sum:", lhs+rhs) // -, *, /
	fmt.Println("Division:", lhs/lhs)
	fmt.Println("MODULUS:", rhs%lhs)

	var unsignInt uint = 200
	//var superLognInt uint64 = 1000000

	fmt.Println("Unsigned sum:", unsignInt*2+uint(200))

	fmt.Println("========================ALIASING=======================")
	var aByte byte
	var aRune rune

	fmt.Printf("%T and %v and %v bytes\n", aByte, aByte, unsafe.Sizeof(aByte))
	fmt.Printf("%T and %v and %v bytes\n", aRune, aRune, unsafe.Sizeof(aRune))

	fmt.Println("===================FLOATING POINT NUMBERS==========================")
	var fFLOAT32 float32 = 23.123
	var fFLOAT64 float64 = 123.123123

	fmt.Println("Floating sum:", float64(fFLOAT32)*fFLOAT64)

	fmt.Println("===================COMPLEX NUMBERS============================")
	c1 := complex(2, 5)
	c2 := 12 + 2i

	c3 := c1 + c2
	fmt.Println("Complex sum:", c3)
	fmt.Println("Complex mult:", c1*c2)

	fmt.Println("=========================STRINGS=============================")
	name := "Bob"
	lastname := "Петров"

	fmt.Println("Concatenate:", name+lastname)
	fmt.Println("Length of name:", len(name))
	fmt.Println("Length of lastname:", len(lastname))
	fmt.Println("Length(runes):", utf8.RuneCountInString(lastname))

	substring := "f"
	generalString := "float"
	fmt.Println(strings.Contains(generalString, substring))

	var LargeUint uint8 = 112
	var LargeInt int8 = int8(LargeUint)
	fmt.Println("OVERFLOW EXAMPLE:", LargeInt)

}
