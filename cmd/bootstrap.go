package bootstrap

import (
	"bank-microservice/internal/infrastructure/database"
	"bank-microservice/internal/repositories"
	"bank-microservice/internal/services"
	"bank-microservice/internal/services/bank"
	"bank-microservice/internal/services/card"
	"bank-microservice/internal/services/person"
	"context"
	"fmt"
)

// Инициализация и запуск сервера

type Bootstrapper struct {
	// Вложенность как в пакетах
	//conltroller *controllers.
	//database
	DB           *database.Database
	Repositories repositories.Repositories
	Services     services.Services
}

func New() (*Bootstrapper, error) {
	var (
		ctx          = context.Background()
		bootstrapper = &Bootstrapper{}
		err          error
	)

	// init
	bootstrapper.DB, err = database.New(ctx)
	if err != nil {
		return nil, err
	}

	bootstrapper.Repositories = repositories.NewRepository(bootstrapper.DB)

	bootstrapper.initService()

	// bootstrap
	return bootstrapper, nil
}

func (b *Bootstrapper) initService() {
	b.Services = services.Services{
		Card:   card.NewService(b.Repositories.Card),
		Bank:   bank.NewService(b.Repositories.Bank),
		Person: person.NewService(b.Repositories.Person),
	}
}

func (b *Bootstrapper) RunAPI() {
	// TODO
	fmt.Println("Bootstrapper.RunAPI()")
}
