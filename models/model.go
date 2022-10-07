package models

type SinglePerson struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Story string `json:"story"`
}

var Persons []SinglePerson
