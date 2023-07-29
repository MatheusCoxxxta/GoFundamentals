package main

import "fmt"

type Person struct {
	Name string
}

type NormalPerson struct {
	Person
	Age uint8
}

func (p NormalPerson) String() string {
	return fmt.Sprintf("Hello, my name is %s and I'm %d", p.Name, p.Age)
}

type LegalPerson struct {
	Person
	Active bool
}

func (p LegalPerson) String() string {
	var status = "active"

	if !p.Active {
		status = "inactive"
	}

	return fmt.Sprintf("Hello, I'm %s and I'm an %s company", p.Name, status)
}

func main() {
	normalPerson := NormalPerson{Age: 21, Person: Person{Name: "Matheus"}}
	fmt.Println(normalPerson)

	legalPerson := LegalPerson{Active: true, Person: Person{Name: "Matheus"}}
	fmt.Println(legalPerson)
}
