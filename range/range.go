package main

import (
	"fmt"
)

func main() {

	var message string = "Hello word"

	for i, v := range message {
		fmt.Printf("Index: %d, value: %c\n", i, v)
	}

}
