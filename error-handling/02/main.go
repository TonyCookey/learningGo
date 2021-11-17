package main

import (
	"log"
	"os"
)

func init() {
	of, err := os.Open("log.txt")
	if err != nil {
		nf, err := os.Create("log.txt")
		if err != nil {
			log.Println(err)
		}
		log.SetOutput(nf)
	}
	log.SetOutput(of)
}
func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		// log.Println() adds a date to the log output compared to fmt.Println which just prints out the error
		log.Println(err)
	}
}
