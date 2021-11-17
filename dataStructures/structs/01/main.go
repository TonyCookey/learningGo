package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
	race      string
	religion  string
}

func main() {
	person1 := person{
		firstName: "Tony",
		lastName:  "Cookey",
		age:       23,
		race:      "Black",
		religion:  "Christianity",
	}
	fmt.Println(person1)
}
