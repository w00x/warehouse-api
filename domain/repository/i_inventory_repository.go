package repository

import (
	"warehouse/domain"
	"warehouse/infrastructure/errors"
)

type IInventoryRepository interface {
	All() (*[]domain.Inventory, errors.IBaseError)
	Find(id string) (*domain.Inventory, errors.IBaseError)
	Create(instance *domain.Inventory) (*domain.Inventory, errors.IBaseError)
	Update(instance *domain.Inventory) (*domain.Inventory, errors.IBaseError)
	Delete(instance *domain.Inventory) errors.IBaseError
}
