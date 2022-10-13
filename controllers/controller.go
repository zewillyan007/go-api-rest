package controllers

import (
	"encoding/json"
	"fmt"
	"go-api-rest/database"
	"go-api-rest/models"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func AllPersons(w http.ResponseWriter, r *http.Request) {

	p := []models.Persons{}
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func ReturnOnePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	p := models.Persons{}
	database.DB.First(&p, id)
	json.NewEncoder(w).Encode(p)
}
