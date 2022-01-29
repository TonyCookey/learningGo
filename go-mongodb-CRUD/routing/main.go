package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"learningGo/go-mongodb-CRUD/routing/models"
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
	user := models.User{
		FirstName: "Anthony",
		LastName:  "Iroegbu",
		Age:       24,
	}
	profile, err := json.Marshal(user)

	_, err = fmt.Fprint(res, string(profile))
	if err != nil {
		log.Fatal(err)
	}
}
