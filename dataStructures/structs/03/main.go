package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
	race      string
	religion  string
	Personality
}

type Personality struct {
	temperament     string
	personalityType string
}

func main() {
	// Inheritance and Promotion of Structs
	person1 := Person{
		firstName: "Tony",
		lastName:  "Cookey",
		age:       23,
		race:      "Black",
		religion:  "Christianity",
		Personality: Personality{
			temperament:     "Sanguine",
			personalityType: "Introverted Extrovert",
		},
	}
	fmt.Println(person1.fullName())
}
func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}
