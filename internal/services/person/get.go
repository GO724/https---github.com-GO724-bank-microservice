package person

import (
	"bank-microservice/internal/entity"
	"context"
	"fmt"
)

func (b *personService) Get(ctx context.Context, inn uint) (*entity.Person, error) {
	fmt.Println("service: get person.inn[%d]", inn)
	return &entity.Person{}, nil
}
