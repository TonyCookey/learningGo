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
	var p1 Person
	bs := []byte(`{"first_name":"Tony","last_name":"Cookey","username":"tonycookey"}`)
	err := json.Unmarshal(bs, &p1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(p1)

}
