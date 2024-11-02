package person

import (
	"bank-microservice/internal/infrastructure/database"
	"fmt"
)

type Person interface {
}

type personRepository struct {
	db *database.Database
}

func NewRepository(db *database.Database) *personRepository {
	fmt.Println("new person repository")
	return &personRepository{
		db: db,
	}
}

//READY:TODO: пофиксить все функции, и вынести их в отдельные файлы одноименные
