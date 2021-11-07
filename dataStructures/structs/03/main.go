package main

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
