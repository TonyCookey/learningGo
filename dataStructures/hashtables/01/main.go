package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// get the book moby dick
	res, err := http.Get("https://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	}

	// scan the page
	// NewScanner takes a reader and res.Body implements the reader interface (so it is a reader)
	var scanner = bufio.NewScanner(res.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)
	// Loop over the words
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
