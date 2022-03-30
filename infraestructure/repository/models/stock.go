package models

import (
	"gorm.io/gorm"
	"warehouse/shared"
)

type Stock struct {
	gorm.Model
	Id				uint		`gorm:"primaryKey;autoIncrement"`
	ItemId			uint
	Item 			*Item
	RackId			uint
	Rack 			*Rack
	Quantity        int
	OperationDate  	shared.DateTime
	ExpirationDate 	shared.DateTime
}

func NewStock(id uint, item *Item, rack *Rack, quantity int, operationDate shared.DateTime,
	expirationDate shared.DateTime, ItemId uint, RackId uint) *Stock {
	return &Stock{Id: id, Item: item, Rack: rack, Quantity: quantity, OperationDate: operationDate,
		ExpirationDate: expirationDate, ItemId: ItemId, RackId: RackId}
}