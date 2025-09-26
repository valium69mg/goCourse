package main

import (
	"fmt"
	"reflect"
)

func main() {
	// initialize map
	mapVariable := make(map[string]int)
	mapVariable["Carlos"] = 98
	fmt.Println(mapVariable)
	// iterate over it
	for key, value := range mapVariable {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
	// delete entry
	delete(mapVariable, "Carlos")
	fmt.Println(mapVariable)
	// check if exists
	var alice string = "Alice"
	mapVariable[alice] = 70
	val, exists := mapVariable["Alice"]
	if exists {
		fmt.Println("Found:", val)
	} else {
		fmt.Println("Not Found:", alice)
	}
	// clear
	clear(mapVariable)
	fmt.Println(mapVariable)
	// blank identifier
	_, valueExists := mapVariable["ximena"]
	if !valueExists {
		fmt.Println("Not Found!")
	}
	mapVariable2 := make(map[string]int)

	if reflect.DeepEqual(mapVariable, mapVariable2) {
		fmt.Println("Maps are equal!")
	}
}
