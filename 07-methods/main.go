package main

import "fmt"

type Person struct {
	Name     string
	Age      uint8
	Status   bool
	document string
}

func (p Person) String() string {
	return fmt.Sprintf("Hello, my name is '%s'", p.Name)
}

type Category struct {
	Name   string
	Parent *Category
}

func (c Category) HasParent() bool {
	return c.Parent != nil
}

func (c *Category) SetParent(parent *Category) {
	c.Parent = parent
}

func main() {
	category := Category{
		Name: "Category I",
	}

	if category.HasParent() {
		fmt.Println("Has parent")
	} else {
		fmt.Println("Has no parent")
	}

	person := Person{
		Name:   "Matheus",
		Age:    21,
		Status: true,
	}

	fmt.Println(person)

	category.SetParent(&Category{Name: "Parent Category"})

	fmt.Println(category.Parent)

	if category.HasParent() {
		fmt.Println("Has parent")
	} else {
		fmt.Println("Has no parent")
	}
}
