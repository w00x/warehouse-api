package repository

import (
	"warehouse/domain"
	"warehouse/infraestructure/errors"
)

type IInventoryRepository interface {
	All() (*[]domain.Inventory, errors.IBaseError)
	Find(id uint) (*domain.Inventory, errors.IBaseError)
	Create(instance *domain.Inventory) (*domain.Inventory, errors.IBaseError)
	Update(instance *domain.Inventory) (*domain.Inventory, errors.IBaseError)
	Delete(instance *domain.Inventory) errors.IBaseError
}
