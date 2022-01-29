package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"learningGo/go-mongodb-CRUD/mvc/controllers"
	"log"
	"net/http"
)

func main() {
	router := httprouter.New()
	userController := controllers.NewUserController(getSession())
	router.GET("/", index)
	router.GET("/user/:id", userController.Profile)
	router.POST("/create", userController.CreateUser)
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
func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		log.Fatal(err)
	}
	return session
}
