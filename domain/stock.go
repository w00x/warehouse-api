package domain

import (
	"time"
)

type Stock struct {
	Id				string
	Item 			*Item
	Rack 			*Rack
	Quantity        int
	OperationDate  	time.Time
	ExpirationDate 	time.Time
}

func NewStock(id string, item *Item, rack *Rack, quantity int, operationDate time.Time,
	expirationDate time.Time) *Stock {
	return &Stock{Id: id, Item: item, Rack: rack, Quantity: quantity, OperationDate: operationDate,
		ExpirationDate: expirationDate}
}
