package main

import "fmt"

// package scope variable
var (
	firstName string
	phone     string
	Address   string // accessible for other packages
)

func main() {
	// function scope variable
	var name string
	name = "Matheus"

	// := is the same as declare and evaluate a variable
	message := ":= is the same as declare and evaluate a variable"

	fmt.Println("Hello ", name, "!")
	fmt.Println("Message: ", message)

	// Changing the variable value
	var x, y = 10, 20
	fmt.Println(x, y)

	x, y = y, x
	fmt.Println(x, y)
}
