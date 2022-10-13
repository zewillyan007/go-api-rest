package routes

import (
	"go-api-rest/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleResquest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/persons", controllers.AllPersons).Methods("GET")
	r.HandleFunc("/api/persons", controllers.CreateNewPerson).Methods("POST")
	r.HandleFunc("/api/persons/{id}", controllers.ReturnOnePerson).Methods("GET")
	r.HandleFunc("/api/persons/{id}", controllers.DeletePerson).Methods("DELETE")
	r.HandleFunc("/api/persons/{id}", controllers.EditPerson).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", r))
}
