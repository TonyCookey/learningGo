package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type Person struct {
	First string
	Last  string
	Age   int
}

func main() {
	var p1 Person
	reader := strings.NewReader(`{
		"First": "Tony",
		"Last":  "Cookey",
		"Age":   23
	}`)
	//prints to standard out stdout
	err := json.NewDecoder(reader).Decode(&p1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(p1)
}