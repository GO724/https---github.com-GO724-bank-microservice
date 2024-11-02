package card

import (
	"bank-microservice/internal/entity"
	"context"
	"fmt"
	"time"
)

func (c cardRepository) Get(ctx context.Context, id uint) (entity.Card, error) {
	// Read record
	// sqlQuery := "SELECT bank.name FROM bank WHERE bank.bic=%1"

	// var name string

	//err := b.db..db.QueryRow(ctx, sqlQuery, bic).Scan(&name)
	//	if err != nil {
	//		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	//	}
	//
	//	var bank entity.Bank
	//	bank.Bic = bic
	//	bank.Name = name
	//	return bank, err
	fmt.Println("service: get card.id[%d]\n", id)
	return entity.Card{Id: 123, Person: entity.Person{Inn: 456, Name: "testPersonName"}, Bank: entity.Bank{Bic: 789, Name: "nestBankName"}, Expires: time.Now()}, nil
}
