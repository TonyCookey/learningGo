package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string
	LastName  string
	Username  string
	password  string
	Age       int
}

func main() {
	p1 := Person{
		FirstName: "Tony",
		LastName:  "Cookey",
		Username:  "tonycookey",
		// unexported field - started with lower case
		password: "!password",
		Age:      23,
	}
	// only Exported fields will be Marshalled
	bs, err := json.Marshal(p1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T\n", bs)
	fmt.Println(string(bs))
}
