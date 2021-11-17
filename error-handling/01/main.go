package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		// log.Println() adds a date to the log output compared to fmt.Println which just prints out the error
		log.Println(err)
	}
}
