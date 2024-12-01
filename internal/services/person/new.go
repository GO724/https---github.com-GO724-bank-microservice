package person

import (
	"bank-microservice/internal/entity"
	"context"
	"fmt"
)

func (b *personService) New(ctx context.Context, person entity.Person) (*entity.Person, error) {
	fmt.Println("service: new person[%v]", person)
	return &entity.Person{}, nil
}
