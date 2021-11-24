package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}
type student struct {
	person
	course string
	level  string
}
type human interface {
	fullName()
}

func (s student) fullName() {
	fmt.Println(s.firstName, s.lastName)
}
func (p person) fullName() {
	fmt.Println("Person:", p.firstName, p.lastName)
}
func (s student) identity() {
	fmt.Println("Student:", s.course, s.level)
}
func saySomething(h human) {
	h.fullName()
}
func main() {
	p1 := person{
		firstName: "Tony",
		lastName:  "Cookey",
	}
	s1 := student{
		person: person{
			firstName: "Clara",
			lastName:  "Johnson",
		},
		course: "Humanities",
		level:  "400",
	}
	p1.fullName()
	s1.fullName()
	s1.person.fullName()
	s1.identity()

	// using interface
	saySomething(s1)
	saySomething(p1)
}
