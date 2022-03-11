package domain

import (
	"time"
)

type Stock struct {
	Id				uint
	ItemId			uint
	Item 			*Item
	RackId			uint
	Rack 			*Rack
	Quantity        int
	OperationDate  	time.Time
	ExpirationDate 	time.Time
}

func NewStock(id uint, item *Item, rack *Rack, quantity int, operationDate time.Time,
	expirationDate time.Time, ItemId uint, RackId uint) *Stock {
	return &Stock{Id: id, Item: item, Rack: rack, Quantity: quantity, OperationDate: operationDate,
		ExpirationDate: expirationDate, ItemId: ItemId, RackId: RackId}
}
