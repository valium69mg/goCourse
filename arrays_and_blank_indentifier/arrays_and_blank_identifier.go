package main

import "fmt"

func main() {
	// INTEGER
	var numbers [10]int
	numbers[0] = 1 // rest of values initialize in 0
	fmt.Println(numbers)

	// STRING
	fruits := [3]string{"apple", "orange", "pear"}
	fmt.Println(fruits)

	// COPY AN ARRAY
	original := [1]int{1}
	copy := original
	original[0] = 0
	fmt.Println(original)
	fmt.Println(copy)

	// ITERATION
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("ELement at index, %d, value, %d\n", i, numbers[i])
	}
	// MODERNIZED ITERATION
	for index, value := range numbers {
		fmt.Printf("ELement at index, %d, value, %d\n", index, value)
	}

	// Blank identifier
	for _, v := range fruits {
		fmt.Println("Fruit " + v)
	}

	// matrix
	var matrix [2][2]int = [2][2]int{
		{2, 2},
		{2, 2},
	}

	fmt.Println(matrix)

	// POINTERS
	var ptr *[2][2]int = &matrix
	var copiedMatrix [2][2]int = *ptr
	copiedMatrix[0][0] = 1
	fmt.Println(copiedMatrix)

}
