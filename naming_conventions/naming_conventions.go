package main

import "fmt"

type EmployeeGoogle struct {
	FirstName string
	LastName  string
	Age       int
}

type EmployeeApple struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// PascalCase: strucs, interfaces, enums

	// snake_case: variables, constants and file names

	// UPPERCASE: constants

	// camelCase : js, html, other libraries

	const MAXRETRIES = 5

	var employeeID int
	employeeID = 1001
	fmt.Println("Emplooye ID: ", employeeID)

}
