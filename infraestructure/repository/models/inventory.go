package models

import (
	"gorm.io/gorm"
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