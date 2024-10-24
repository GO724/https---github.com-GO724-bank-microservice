package person

import (
	"bank-microservice/internal/entity"
	"context"
)

func (p personRepository) New(ctx context.Context, person entity.Person) error {
	// New record
	// err := p.db.ExecSql(ctx, fmt.Sprintf("INSERT INTO person(inn,name) VALUES (%d,%s)", person.Inn, person.Name))
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "unable to create person[%v] : %v\n", person, err)
	// } else {
	// 	fmt.Fprintf(os.Stderr, "create person[%v]\n", person)
	// }
	//return err
	return nil
}
