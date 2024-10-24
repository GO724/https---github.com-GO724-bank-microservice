package bank

import (
	"bank-microservice/internal/entity"
	"context"
)

func (b bankRepository) Get(ctx context.Context, bic uint) (entity.Bank, error) {
	// Read record
	//sqlQuery := "SELECT bank.name FROM bank WHERE bank.bic=%1"

	// var name string

	// err := b.db..db.QueryRow(ctx, sqlQuery, bic).Scan(&name)
	//	if err != nil {
	//		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	//	}
	//
	//	var bank entity.Bank
	//	bank.Bic = bic
	//	bank.Name = name
	//	return bank, err
	return entity.Bank{Bic: 0, Name: "0"}, nil
}
