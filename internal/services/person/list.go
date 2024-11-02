package person

import (
	"bank-microservice/internal/entity"
	"context"
	"fmt"
)

func (b *personService) List(ctx context.Context, person entity.Person) ([]*entity.Person, error) {
	fmt.Println("service: list person[%v]", person)
	return make([]*entity.Person, 0), nil
}
