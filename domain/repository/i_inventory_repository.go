package repository

import (
	"time"
	"warehouse/domain"
	"warehouse/infraestructure/errors"
)

type IInventoryRepository interface {
	All() ([]*domain.Inventory, errors.IBaseError)
	Find(id string) (*domain.Inventory, errors.IBaseError)
	Update(id string, operationDate time.Time) errors.IBaseError
	Create(operationDate time.Time) (*domain.Inventory, errors.IBaseError)
	Delete(id string) errors.IBaseError
}
