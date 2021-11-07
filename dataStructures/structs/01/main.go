package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
	race      string
	religion  string
}

func main() {
	person1 := Person{
		firstName: "Tony",
		lastName:  "Cookey",
		age:       23,
		race:      "Black",
		religion:  "Christianity",
	}
	fmt.Println(person1)
}
