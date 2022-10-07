package controllers

import (
	"encoding/json"
	"fmt"
	"go-api-rest/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func AllPersons(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Persons)
}

func ReturnOnePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, person := range models.Persons {
		if strconv.Itoa(person.Id) == id {
			json.NewEncoder(w).Encode(person)
		}
	}
}
