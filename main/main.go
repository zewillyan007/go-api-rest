package main

import (
	"fmt"
	"go-api-rest/database"
	"go-api-rest/routes"
)

func main() {
	database.ConnectDatabase()

	fmt.Println("Starting Rest Server With Go")
	routes.HandleResquest()
}
