package main

import "fmt"

func main() {

	names := []string{"Javascripto", "Golango", "Elixar"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for _, name := range names {
		fmt.Println(name)
	}

	var i int

	for i < len(names) {
		fmt.Println(names[i])
		i++
	}

	// infinite loop
	for {
		fmt.Println(names[i])
		i++
	}

}
