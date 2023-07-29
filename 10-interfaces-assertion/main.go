package main

import "fmt"

type Document interface {
	Doc() string
}

type Person struct {
	Name string
}

type NormalPerson struct {
	Person
	Age uint8
	Cpf string
}

func (np NormalPerson) String() string {
	return fmt.Sprintf("Hello, my name is %s and I'm %d", np.Name, np.Age)
}

func (np NormalPerson) Doc() string {
	return np.Cpf
}

type LegalPerson struct {
	Person
	Active bool
	Cnpj   string
}

func (lp LegalPerson) String() string {
	var status = "active"

	if !lp.Active {
		status = "inactive"
	}

	return fmt.Sprintf("Hello, I'm %s and I'm an %s company. My document: %s", lp.Name, status, lp.Doc())
}

func (lp LegalPerson) Doc() string {
	return lp.Cnpj
}

func show(d Document) {
	fmt.Println(d)

	switch d.(type) {
	case NormalPerson:
		fmt.Println("CPF:", d.(NormalPerson).Cpf)
	case LegalPerson:
		fmt.Println("CNPJ:", d.(LegalPerson).Cnpj)
	default:
		fmt.Println("Unknown document type")
	}

	// Other option

	switch localDocument := d.(type) {
	case NormalPerson:
		fmt.Println("CPF:", localDocument.Doc())
	case LegalPerson:
		fmt.Println("CNPJ:", localDocument.Doc())
	default:
		fmt.Println("Unknown document type")
	}

}

func main() {
	normalPerson := NormalPerson{Age: 21, Cpf: "4321", Person: Person{Name: "Matheus"}}
	show(normalPerson)

	legalPerson := LegalPerson{Active: true, Cnpj: "1234", Person: Person{Name: "Matheus"}}
	show(legalPerson)
}
