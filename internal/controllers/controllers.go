package controllers

import (
	"bank-microservice/internal/services/bank"
	"bank-microservice/internal/services/card"
	"bank-microservice/internal/services/person"
)

type controller struct {
	// контроллер знает только о сервисах!
	bank   bank.Bank
	person person.Person
	card   card.Card
}

func New() (Controller *controller) {
	//ctx := context.Background()

	// bs, err := bank.NewService(ctx)
	// // if err != nil {
	// // 	log.Fatalln("can't init obj.bank")
	// // }

	return &controller{
		// bank:   b,
		// person: nil,
		// card:   nil,
	}
}
