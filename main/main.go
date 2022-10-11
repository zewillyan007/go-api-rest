package main

import (
	"fmt"
	"go-api-rest/database"
	"go-api-rest/models"
	"go-api-rest/routes"
)

func main() {
	models.Persons = []models.SinglePerson{
		{Id: 1, Name: "Name 1", Story: "Story 1"},
		{Id: 2, Name: "Name 2", Story: "Story 2"},
		{Id: 3, Name: "Name 3", Story: "Story 3"},
	}

	database.ConnectDatabase()

	fmt.Println("Starting Rest Server With Go")
	routes.HandleResquest()
}
