package person

import (
	"bank-microservice/internal/entity"
	"context"
)

func (p personRepository) Get(ctx context.Context, inn uint) (entity.Person, error) {
	// Read record
	// sqlQuery := "SELECT person.name FROM person WHERE person.inn=%1"

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
	return entity.Person{Inn: '0', Name: "0"}, nil
}
