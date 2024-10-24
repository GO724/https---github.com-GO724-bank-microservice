package bank

import "bank-microservice/internal/repositories/bank"

type Bank interface {
}

type bankService struct {
	bankRepo bank.Bank
}

func NewService(b bank.Bank) Bank {
	return &bankService{
		bankRepo: b,
	}
}
