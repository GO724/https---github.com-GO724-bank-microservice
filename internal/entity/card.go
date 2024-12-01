package entity

import (
	"time"
)

type Card struct {
	Id      uint
	Person  Person
	Bank    Bank
	Expires time.Time // UTC
}

func (c Card) IsValid() bool {
	return c.Expires.Before(time.Now().UTC())
}
