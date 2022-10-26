package domain

import (
	"github.com/google/uuid"
	"warehouse/shared"
)

type Inventory struct {
	id            string
	OperationDate shared.DateTime
}

func NewInventory(id string, operationDate shared.DateTime) *Inventory {
	if id == "" {
		id = uuid.New().String()
	}
	return &Inventory{id: id, OperationDate: operationDate}
}

func (i Inventory) Id() string {
	return i.id
}
