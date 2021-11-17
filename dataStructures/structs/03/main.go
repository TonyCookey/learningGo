package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
	race      string
	religion  string
	personality
}

type personality struct {
	temperament     string
	personalityType string
}

func main() {
	// Inheritance and Promotion of Structs
	person1 := person{
		firstName: "Tony",
		lastName:  "Cookey",
		age:       23,
		race:      "Black",
		religion:  "Christianity",
		personality: personality{
			temperament:     "Sanguine",
			personalityType: "Introverted Extrovert",
		},
	}
	fmt.Println(person1.fullName())
}
func (p person) fullName() string {
	return p.firstName + " " + p.lastName
}
