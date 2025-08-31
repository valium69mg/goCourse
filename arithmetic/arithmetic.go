package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int = 10, 3
	var result int

	result = a + b
	fmt.Println("Addition:", result)
	result = a - b
	fmt.Println("Substraction:", result)
	result = a / b
	fmt.Println("Division:", result)
	result = a * b
	fmt.Println("Multiplication:", result)
	result = a % b
	fmt.Println("Remainder:", result)

	const pi float64 = 3.14
	c := 1.0
	var floatResult float64 = c / pi
	fmt.Println("Division:", floatResult)

	var maxInt int64 = 9223372036854775807
	fmt.Println(maxInt + 1) // overflow

	var uMaxInt uint64 = 18446744073709551615 // overflow unsigned int
	uMaxInt = uMaxInt + 1
	fmt.Println(uMaxInt)

	var smallFLoat float64 = 1.0e-323
	smallFLoat = smallFLoat / math.MaxFloat64
	fmt.Println(smallFLoat)

}
