package models

import (
	"gorm.io/gorm"
	"warehouse/domain"
	"warehouse/shared"
)

type Inventory struct {
	gorm.Model
	Id 				uint		`gorm:"primaryKey;autoIncrement"`
	OperationDate 	shared.DateTime
}

func NewInventory(id uint, operationDate shared.DateTime) *Inventory {
	return &Inventory{Id: id, OperationDate: operationDate}
}

func FromInventoryDomainToModel(i *domain.Inventory) *Inventory {
	return NewInventory(i.Id, i.OperationDate)
}

func (i Inventory) ToDomain() *domain.Inventory {
	return domain.NewInventory(i.Id, i.OperationDate)
}