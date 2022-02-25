package repository

import (
	"time"
	"warehouse/domain"
	"warehouse/infraestructure/errors"
)

type IInventoryRepository interface {
	All() ([]*domain.Inventory, error)
	Find(id string) (*domain.Inventory, error)
	Update(id string, operationDate time.Time) errors.BaseError
	Create(operationDate time.Time) (*domain.Inventory, error)
	Delete(id string) error
}
