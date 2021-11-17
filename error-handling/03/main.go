package main

import (
	"log"
	"os"
)

func init() {
	nf, err := os.Create("log.txt")
	if err != nil {
		log.Println(err)
	}
	log.SetOutput(nf)
}
func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		// log.Println() adds a date to the log output compared to fmt.Println which just prints out the error
		log.Println(err)
	}
}
