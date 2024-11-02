package person

import (
	"bank-microservice/internal/entity"
	"context"
	"fmt"
)

func (p personRepository) List(ctx context.Context, person entity.Person) ([]entity.Person, error) {
	//	// List of record filtered on p(Person)
	//	sb := make([]entity.Person, 0, 10)
	//	return sb, nil
	fmt.Println("service: list all of person[%v]", person)
	return nil, nil
}
