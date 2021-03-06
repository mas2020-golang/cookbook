package main

import "fmt"

/*
This example shows how to use the type composition to simulate inheritance.
*/

type BaseType struct {
	Name    string
	Surname string
	Age     int
}

func (b *BaseType) Greetings() {
	fmt.Println("Hello from BaseType!")
	(*b).Name = "test"
}

/*
Specific type contains the BaseType struct
*/
type SpecificType struct {
	*BaseType // using a pointer is better for performance reason
	Type      string
	Category  string
}

func createSpecificType() {
	// new SpecificType using literal
	st1 := SpecificType{
		Type:     "Type of spec type",
		Category: "Cat of spec type",
		BaseType: &BaseType{Name: "My name", Surname: "Surname from Base Type", Age: 12},
	}
	// IMPORTANT: you can directly use the methods of the BaseType structure implicit included in the SpecificType
	st1.Greetings()
	fmt.Println("value of st1.BaseType is", st1.BaseType)
}
