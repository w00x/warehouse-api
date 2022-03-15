package domain

import (
	"warehouse/shared"
)

type Inventory struct {
	Id 				uint
	OperationDate 	shared.DateTime
}

func NewInventory(id uint, operationDate shared.DateTime) *Inventory {
	return &Inventory{Id: id, OperationDate: operationDate}
}
