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

func (uc *UserController) Profile(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	fmt.Println(id)

	if !bson.IsObjectIdHex(id) {
		res.WriteHeader(http.StatusNotFound)
		return
	}
	ID := bson.ObjectIdHex(id)

	user := models.User{}

	err := uc.session.DB("learning-go").C("users").FindId(ID).One(&user)
	if err != nil {
		fmt.Println("Could not find user", err)
		res.WriteHeader(http.StatusNotFound)
		return
	}
	result, err := json.Marshal(&user)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	fmt.Fprint(res, string(result))
}

func (uc *UserController) CreateUser(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	newUser := models.User{}

	//decode from the request body stream
	json.NewDecoder(req.Body).Decode(&newUser)

	// generate new user ID
	newUser.ID = bson.NewObjectId()
	fmt.Println(newUser.ID)

	// store the new user in the DB
	err := uc.session.DB("learning-go").C("users").Insert(newUser)
	if err != nil {
		fmt.Println("could not insert")
		res.WriteHeader(http.StatusNotFound)
		return
	}

	result, err := json.Marshal(newUser)
	if err != nil {
		log.Fatal(err)
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	fmt.Fprint(res, string(result))
}
