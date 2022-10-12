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

	// db := database.ConnectDatabase()
	// allPersons, err := db.Query("select * from single_people sp")
	// if err != nil {
	// 	panic(err.Error())
	// }

	// p := models.SinglePerson{}

	// for allPersons.Next() {
	// 	var id int
	// 	var nameperson, story string

	// 	err := allPersons.Scan(&id, &nameperson, &story)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}

	// 	p.Id = id
	// 	p.NamePerson = nameperson
	// 	p.Story = story

}

func ReturnOnePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	p := models.Persons{}
	database.DB.First(&p, id)
	json.NewEncoder(w).Encode(p)
}
