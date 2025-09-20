package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {

	min := 0
	max := 10
	var randInt int = getRandomInt(min, max)

	// start game
	fmt.Println("Let's play higher or lower!")
	for {
		fmt.Printf("guess the number (%d to %d):", min, max)
		var inputInt int
		_, err := fmt.Scanln(&inputInt)
		if err != nil {
			log.Fatal("Oops, something went wrong: ", err)
		}
		if inputInt > randInt {
			fmt.Println("Lower!")
		} else if inputInt < randInt {
			fmt.Println("Higher!")
		} else {
			fmt.Printf("You guessed correctly, the number is %d\n", randInt)
			break
		}
	}
}

func getRandomInt(min, max int) int {
	if min > max {
		panic("min should not be greater than max")
	}
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	return random.Intn(max-min+1) + min
}
