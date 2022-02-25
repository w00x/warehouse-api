package domain

import (
	"time"
)

type Inventory struct {
	Id 				string
	OperationDate 	time.Time
}

func NewInventory(id string, operationDate time.Time) *Inventory {
	return &Inventory{Id: id, OperationDate: operationDate}
}
