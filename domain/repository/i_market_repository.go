package repository

import (
	"warehouse/domain"
	"warehouse/infrastructure/errors"
)

type IMarketRepository interface {
	All() (*[]domain.Market, errors.IBaseError)
	Find(id string) (*domain.Market, errors.IBaseError)
	Create(instance *domain.Market) (*domain.Market, errors.IBaseError)
	Update(instance *domain.Market) (*domain.Market, errors.IBaseError)
	Delete(instance *domain.Market) errors.IBaseError
}
