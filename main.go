// Сервис для хранения пользовательских реквизитов карточек
// 1. Придумать формат БД(помнить про нормализацию)
// 2. Сделать апи удаление/создание/обновление реквизитов и пользователей
// 3. Сделать асинхронный слушатель, который будет отслеживать реквизиты, и когда они просрочатся удалять их асинхронно

// Каждый пользователь может иметь неограниченное количество платежных карт
// Таблица пользователей
// Таблица платежных карт
// Таблица банков эмитентов платежных карт

// api
// Добавление/Обновление/Удаление/Чтение пользователя
// Добавление/Обновление/Удаление/Чтение банка
// Добавление/Обновление/Удаление/Чтение платежной карты
// Асинхронное удаление просроченных платежных карт

// Реализация:
// Подключение БД при старте. Если нет - возвращать ошибку и на выход? Инициализация только по требованию (вызов метода?)
// api:Card
// IsValid(card_id,person_inn,validDate)(valid bool, err error)
// CreateOrUpdateCard(id,person,bank,validDate)(isNew bool, err error) для err==nill в БД должны быть person bank validDade>Now()
// GetCard(id)(c Card, err error)

// api:Bank
// b.name=""
// b.bic=123
// func (b *bank) Set()(err error)
// func (b *bank) Get()(err error)

package main

import (
	pg "bank-microservice/internal/database/dbEngine/postgres"
	"context"
	"fmt"
	"log"
)

func main() {

	// развертывание сервиса:
	// инициализация БД : чтение параметров БД из конфига (если открытие БД == nil тогда создание новой БД и дефолтного конфига?)

	// cmd/bootstrap
	ctx := context.Background()
	db, err := pg.NewPG(ctx)
	if err != nil {
		log.Fatalln("error connect pg %w", err)
	}

	checkMsg := "check health postgree connect : "
	err = db.Ping(ctx)
	if err != nil {
		fmt.Println(checkMsg + "FAILED")
	} else {
		fmt.Println(checkMsg + "ok")
	}
}
