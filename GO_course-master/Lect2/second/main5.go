package main

import (
	"fmt"
	"math"
)

func main() {
	var a = math.Sqrt(16)
	//const b = math.Sqrt(4)
	const c = "Sam"

	var zeroInt int
	const d int64 = 0 // non-default for const

	result := d + int64(12)

	//typed const
	const aConst int = 22
	var intTemp8 int8 = int8(aConst)
	var intTemp16 int16 = int16(aConst)
	var floatTemp32 float32 = float32(aConst)
	var complexTemp64 complex64 = complex64(aConst)
	fmt.Println(a, c, zeroInt, d, result)
	fmt.Println(intTemp8, intTemp16, floatTemp32, complexTemp64)

	//untyped const
	const bConst = 34
	var inttTemp8 int8 = bConst
	var inttTemp32 int32 = bConst
	var floattTemp32 float32 = bConst
	var complexxTemp64 complex64 = bConst
	fmt.Println(inttTemp8, inttTemp32, floattTemp32, complexxTemp64)
}
