package domain

import (
	"github.com/google/uuid"
	"warehouse/shared"
)

type Stock struct {
	id             string
	ItemId         string
	Item           *Item
	RackId         string
	Rack           *Rack
	Quantity       int
	Comment        string
	OperationDate  shared.DateTime
	ExpirationDate shared.DateTime
}

func NewStock(id string, item *Item, rack *Rack, quantity int, operationDate shared.DateTime,
	expirationDate shared.DateTime, comment string, ItemId string, RackId string) *Stock {
	if id == "" {
		id = uuid.New().String()
	}

	return &Stock{id: id, Item: item, Rack: rack, Quantity: quantity, OperationDate: operationDate,
		ExpirationDate: expirationDate, Comment: comment, ItemId: ItemId, RackId: RackId}
}

func (i Stock) Id() string {
	return i.id
}
