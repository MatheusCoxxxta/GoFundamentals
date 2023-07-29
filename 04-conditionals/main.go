package main

import (
	"fmt"
	"log"
	"os"
)

var (
	heads, tails int
)

func flipCoin(side string) {
	switch side {
	case "heads":
		heads++
		// Go does not use 'break' before condition
	case "tails":
		tails++

	default:
		fmt.Println(("Invalid option"))
	}
}

func main() {
	a, b := 10, 15

	if a > b {
		fmt.Println("a is bigger than b")
	} else if a < b {
		fmt.Println("b is bigger than a")
	} else {
		fmt.Println("a = b")
	}

	// Go supports conditions on case
	switch {
	case a > b:
		fmt.Println("a is bigger than b")
	case a < b:
		fmt.Println("b is bigger than a")
	default:
		fmt.Println("a = b")
	}

	// function scope err
	file, err := os.Open("hello.txt")

	if err != nil {
		log.Panic(err)
	}

	data := make([]byte, 100)

	// local scope err (block = if)
	if _, err := file.Read(data); err != nil {
		log.Panic(err)
	}

	fmt.Println(string(data))
}
