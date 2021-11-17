package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		// calls os.exit - add an exit status 1
		//log.Fatal(err)
		log.Fatalln(err)
	}
}
