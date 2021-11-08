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
	var p1 Person
	bs := []byte(`{
		"FirstName": "Tony",
		"LastName":  "Cookey",
		"Username":  "tonycookey",
		"password":  "!password",
		"Age":       23
	}`)
	err := json.Unmarshal(bs, &p1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(p1)
}
