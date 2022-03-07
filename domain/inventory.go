package domain

import (
	"time"
)

type Inventory struct {
	Id 				uint
	OperationDate 	time.Time
	CreatedAt    	time.Time
	UpdatedAt    	time.Time
}

func NewInventory(id uint, operationDate time.Time) *Inventory {
	return &Inventory{Id: id, OperationDate: operationDate}
}
