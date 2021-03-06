package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.GET("/", index)
	router.GET("/profile", profile)
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}

func index(res http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	_, err := fmt.Fprint(res, "Welcome to my Server")
	if err != nil {
		log.Fatal(err)
	}
}

func profile(res http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	_, err := fmt.Fprint(res, "Welcome to my Profile")
	if err != nil {
		log.Fatal(err)
	}
}
