package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"learningGo/go-mongodb-CRUD/mvc/models"
	"log"
	"net/http"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(session *mgo.Session) *UserController {
	return &UserController{session}
}

func (uc *UserController) Profile(res http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	_, err := fmt.Fprint(res, "My Profile")
	if err != nil {
		log.Fatal(err)
	}
}

func (uc *UserController) CreateUser(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	newUser := models.User{}

	//decode from the request body stream
	json.NewDecoder(req.Body).Decode(&newUser)

	// generate new user ID
	newUser.ID = bson.NewObjectId()

	// store the new user in the DB
	uc.session.DB("learning-go").C("users").Insert(newUser)

	result, err := json.Marshal(newUser)
	if err != nil {
		log.Fatal(err)
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	fmt.Fprint(res, string(result))
}
