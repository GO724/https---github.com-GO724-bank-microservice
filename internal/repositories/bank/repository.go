package bank

import (
	"bank-microservice/internal/infrastructure/database"
	"fmt"
)

type Bank interface {
}

type bankRepository struct {
	db *database.Database
}

func NewRepository(db *database.Database) *bankRepository {
	fmt.Println("new bank repository")
	return &bankRepository{
		db: db,
	}
}

//TODO: пофиксить все функции, и вынести их в отдельные файлы одноименные
