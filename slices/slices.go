package main

import (
	"fmt"
	"reflect"
)

func main() {
	// simple initialization
	var numbers []int
	numbers = append(numbers, 123)
	fmt.Println(numbers)

	// slice from array
	array := [5]int{1, 2, 3, 4, 5}
	slice := array[0:]
	fmt.Println(array)
	fmt.Println(slice)

	// slice copy
	slice1 := make([]int, len(slice))
	copy(slice1, slice)
	fmt.Println(slice1)

	// iterate
	for i, v := range slice {
		fmt.Println("Index ", i, "Value ", v)
	}

	if reflect.DeepEqual(slice, slice1) {
		fmt.Println("Slice equals slice1")
	}

	// 2D
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println(twoD)
}
