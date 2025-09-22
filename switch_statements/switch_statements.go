package main

import "fmt"

func main() {
	fruit := "apple"

	switch fruit {
	case "apple":
		fmt.Println("It's an apple!")
	case "banana":
		fmt.Println("It's an apple!")
	default:
		fmt.Println("Unknown fruit!")
	}

	day := "Monday"
	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It's weekday!")
	case "Saturday", "Sunday":
		fmt.Println("It's weekend")
	default:
		fmt.Println("Invalid day.")
	}

	number := 15

	switch {
	case number < 10:
		fmt.Println("Lower than 10!")
		fallthrough
	case number >= 10 && number <= 20:
		fmt.Println("Between 10 and 20")
		fallthrough
	case number == 15:
		fmt.Println("Number is 15")
	default:
		fmt.Println("Idk man!")
	}

}
