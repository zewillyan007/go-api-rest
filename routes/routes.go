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
	r.HandleFunc("/api/persons/{id}", controllers.ReturnOnePerson).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}
