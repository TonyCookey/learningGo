package main

import (
	"encoding/json"
	"log"
	"os"
)

type Person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := Person{
		First: "Tony",
		Last:  "Cookey",
		Age:   23,
	}
	//prints to standard out stdout
	err := json.NewEncoder(os.Stdout).Encode(p1)
	if err != nil {
		log.Fatal(err)
	}
}
