package controllers

import (
	db "bank-microservice/internal/database"
)

type controller struct {
	Bank   db.Bank
	Person db.Person
	Card   db.Card
}

func New() (Controller *controller) {
	return &controller{
		Bank:   db.Bank,
		Person: db.Person,
		Card:   db.Card,
	}
}
