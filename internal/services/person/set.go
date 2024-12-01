package person

import (
	"bank-microservice/internal/entity"
	"context"
	"fmt"
)

func (b *personService) Set(ctx context.Context, person entity.Person) (*entity.Person, error) {
	fmt.Println("service: set person[%v]", person)
	return &entity.Person{}, nil
}
