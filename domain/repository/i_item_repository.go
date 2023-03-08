package repository

import (
	"warehouse/domain"
	"warehouse/infrastructure/errors"
)

type IItemRepository interface {
	All() (*[]domain.Item, errors.IBaseError)
	Find(id string) (*domain.Item, errors.IBaseError)
	FindByCode(code string) (*domain.Item, errors.IBaseError)
	Create(instance *domain.Item) (*domain.Item, errors.IBaseError)
	Update(instance *domain.Item) (*domain.Item, errors.IBaseError)
	Delete(instance *domain.Item) errors.IBaseError
}
