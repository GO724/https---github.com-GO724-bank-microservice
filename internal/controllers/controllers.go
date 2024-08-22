package controllers

import db "bdms/internal/database"

type controller struct {
	Bank   db.Bank
	Person db.Person
	Card   db.Card
}

// func New() (Controller controller) {
//  	return &controller{
//  		Bank:		db.Bank,
// 		Person:		db.Person,
// 		Card:		db.Card,
//    	}
// }
