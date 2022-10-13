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

func CreateNewPerson(w http.ResponseWriter, r *http.Request) {
	var NewPersons models.Persons
	json.NewDecoder(r.Body).Decode(&NewPersons)
	database.DB.Create(&NewPersons)
	json.NewEncoder(w).Encode(NewPersons)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var OldPerson models.Persons
	database.DB.Delete(&OldPerson, id)
	json.NewEncoder(w).Encode(OldPerson)
}

func EditPerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var FailPerson models.Persons
	database.DB.First(&FailPerson, id)
	json.NewDecoder(r.Body).Decode(&FailPerson)
	database.DB.Save(&FailPerson)
	json.NewEncoder(w).Encode(FailPerson)
}
