package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Username  string `json:"username,omitempty"`
	Password  string `json:"-"`
}

func main() {
	p1 := Person{
		FirstName: "Tony",
		LastName:  "Cookey",
		Username:  "tonycookey",
		Password:  "!password",
	}
	bs, err := json.Marshal(p1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bs))
	// {"first_name":"Tony","last_name":"Cookey","username":"tonycookey"}
	// it excluded password and used the tags I specified when defining the struct

}
