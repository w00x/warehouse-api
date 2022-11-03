package repository

import (
	"warehouse/domain"
	"warehouse/infrastructure/errors"
)

type IStockRepository interface {
	All() (*[]domain.Stock, errors.IBaseError)
	Find(id string) (*domain.Stock, errors.IBaseError)
	Create(instance *domain.Stock) (*domain.Stock, errors.IBaseError)
	Update(instance *domain.Stock) (*domain.Stock, errors.IBaseError)
	Delete(instance *domain.Stock) errors.IBaseError
	AllByInventory(inventoryId string) (*[]domain.Stock, errors.IBaseError)
}
