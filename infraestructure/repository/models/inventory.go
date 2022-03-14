package models

import (
	"gorm.io/gorm"
	"time"
	"warehouse/domain"
)

type Inventory struct {
	gorm.Model
	Id 				uint		`gorm:"primaryKey;autoIncrement"`
	OperationDate 	time.Time
}

func NewInventory(id uint, operationDate time.Time) *Inventory {
	return &Inventory{Id: id, OperationDate: operationDate}
}

func FromInventoryDomainToModel(i *domain.Inventory) *Inventory {
	return NewInventory(i.Id, i.OperationDate)
}

func (i Inventory) ToDomain() *domain.Inventory {
	return domain.NewInventory(i.Id, i.OperationDate)
}