package main

import "fmt"

func main() {

	// iteration over a range
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// iterate over a collection
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Value %d, index %d\n", value, index)
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Odd number found:", i)
		if i == 5 {
			break
		}
	}

	var rows int = 5
	for i := 1; i <= rows; i++ {
		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	gridX := 20
	gridY := 20
	centerX := gridX / 2.0
	centerY := gridY / 2.0
	radius := (gridX / 2.0) - 1
	for i := range gridX {
		for j := range gridY {
			pointX := i
			pointY := j
			formula := ((pointX - centerX) * (pointX - centerX)) + ((pointY - centerY) * (pointY - centerY))
			rSquared := radius * radius
			if formula < rSquared {
				print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println("")
	}
}
