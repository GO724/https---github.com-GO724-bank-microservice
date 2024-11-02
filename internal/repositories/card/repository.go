package card

import (
	"bank-microservice/internal/infrastructure/database"
	"fmt"
)

type Card interface {
}

type cardRepository struct {
	db *database.Database
}

func NewRepository(db *database.Database) *cardRepository {
	fmt.Println("new card repository")
	return &cardRepository{
		db: db,
	}
}
