package card

import (
	"bank-microservice/internal/entity"
	"context"
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
	return entity.Card{Id: '0', Person: entity.Person{Inn: '0', Name: "0"}, Bank: entity.Bank{Bic: '0', Name: "0"}, Expires: time.Now()}, nil
}
