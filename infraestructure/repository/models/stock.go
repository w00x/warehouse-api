package models

import (
	"gorm.io/gorm"
	"time"
	"warehouse/domain"
)

type Stock struct {
	gorm.Model
	Id				uint		`gorm:"primaryKey;autoIncrement"`
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

func FromStockDomainToModel(i *domain.Stock) *Stock {
	return NewStock(i.Id, FromItemDomainToModel(i.Item), FromRackDomainToModel(i.Rack), i.Quantity, i.OperationDate, i.ExpirationDate, i.ItemId, i.RackId)
}

func (i Stock) ToDomain() *domain.Stock {
	var rack *domain.Rack
	var item *domain.Item
	if i.Rack != nil {
		rack = i.Rack.ToDomain()
	}
	if i.Item != nil {
		item = i.Item.ToDomain()
	}
	return domain.NewStock(i.Id, item, rack, i.Quantity, i.OperationDate, i.ExpirationDate, i.RackId, i.ItemId)
}

func (i Stock) ToStruct() map[string]interface{} {
	return map[string]interface{}{
		"item_id": i.ItemId,
		"rack_id": i.RackId,
		"quantity": i.Quantity,
		"operation_date": i.OperationDate,
		"expiration_date": i.ExpirationDate,
	}
}