package repository

import (
	"warehouse/domain"
	"warehouse/infrastructure/errors"
)

type IPriceRepository interface {
	All() (*[]domain.Price, errors.IBaseError)
	Find(id string) (*domain.Price, errors.IBaseError)
	Create(instance *domain.Price) (*domain.Price, errors.IBaseError)
	Update(instance *domain.Price) (*domain.Price, errors.IBaseError)
	Delete(instance *domain.Price) errors.IBaseError
}
