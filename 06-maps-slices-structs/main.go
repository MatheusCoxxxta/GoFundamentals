package main

import "fmt"

// Struct
type Person struct {
	Name     string
	Age      uint8
	Status   bool   // public
	document string // private
}

// Using pointer to reference struct from struct
type Category struct {
	Name   string
	Parent *Category
}

func main() {
	// Slice: dynamic
	namesSlice := []string{"Matheus", "Costa"}

	fmt.Println(namesSlice, "len:", len(namesSlice), "cap:", cap(namesSlice))

	namesSlice = append(namesSlice, "Name")
	fmt.Println(namesSlice, "len:", len(namesSlice), "cap:", cap(namesSlice))

	// Array: fixed length
	namesArray := make([]string, 10, 20)
	fmt.Println(namesArray, "len:", len(namesArray), "cap:", cap(namesArray))

	// Map
	ages := make(map[string]uint8)
	ages["Matheus"] = 21

	fmt.Println(ages, "len:", len(ages))
	fmt.Println("Matheus:", ages["Matheus"])

	value, ok := ages["Matheus"]
	fmt.Println(value, ok)

	// Struct
	person := Person{
		Name:     "Matheus",
		Age:      21,
		Status:   true,
		document: "1",
	}

	fmt.Println(person)
	fmt.Println(person.Name)
}
