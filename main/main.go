package main

import (
	"fmt"
	"go-api-rest/models"
	"go-api-rest/routes"
)

func main() {
	models.Persons = []models.SinglePerson{
		{Id: 1, Name: "Name 1", Story: "Story 1"},
		{Id: 2, Name: "Name 2", Story: "Story 2"},
	}

	fmt.Println("Starting Rest Server With Go")
	routes.HandleResquest()
}
