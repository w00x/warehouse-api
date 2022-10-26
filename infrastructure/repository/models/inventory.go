package models

import (
	"gorm.io/gorm"
	"warehouse/shared"
)

type Inventory struct {
	gorm.Model
	ID            string `gorm:"type:uuid;primaryKey;column:id"`
	OperationDate shared.DateTime
}

func NewInventory(id string, operationDate shared.DateTime) *Inventory {
	return &Inventory{ID: id, OperationDate: operationDate}
}
