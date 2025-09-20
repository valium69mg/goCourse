package main

import "fmt"

func main() {
	// for as while loop
	var i int = 1
	for {
		fmt.Printf("Iteration: %d\n", i)
		if i == 10 {
			break
		}
		i++
	}
}
