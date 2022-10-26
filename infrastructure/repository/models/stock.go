package models

import (
	"gorm.io/gorm"
	"warehouse/shared"
)

type Stock struct {
	gorm.Model
	ID             string `gorm:"type:uuid;primaryKey;column:id"`
	ItemId         string
	Item           *Item
	RackId         string
	Rack           *Rack
	Quantity       int
	OperationDate  shared.DateTime
	ExpirationDate shared.DateTime
	Comment        string
}

func NewStock(id string, item *Item, rack *Rack, quantity int, operationDate shared.DateTime,
	expirationDate shared.DateTime, itemId string, rackId string, comment string) *Stock {
	return &Stock{ID: id, Item: item, Rack: rack, Quantity: quantity, OperationDate: operationDate,
		ExpirationDate: expirationDate, ItemId: itemId, RackId: rackId, Comment: comment}
}
