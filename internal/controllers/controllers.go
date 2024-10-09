package controllers

import (
	db "bank-microservice/internal/database"
	pg "bank-microservice/internal/database/dbEngine/postgres"
	"context"
	"log"
)

type controller struct {
	bank   db.Bank
	person db.Person
	card   db.Card
}

func New() (Controller *controller) {
	ctx := context.Background()

	b, err := pg.NewBank(ctx)
	if err != nil {
		log.Fatalln("can't init obj.bank")
	}

	return &controller{
		bank:   b,
		person: nil,
		card:   nil,
	}
}
