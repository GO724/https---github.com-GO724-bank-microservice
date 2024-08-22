package entity

import (
	"time"
)

type Card struct {
	id      uint
	person  Person
	bank    Bank
	expires time.Time // UTC
}

func (c Card) isValid() (result bool) {
	result = c.expires.Before(time.Now().UTC())
	return result
}
