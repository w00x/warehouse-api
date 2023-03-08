package repository

import (
	"warehouse/domain"
	"warehouse/infrastructure/errors"
)

type IRackRepository interface {
	All() (*[]domain.Rack, errors.IBaseError)
	Find(id string) (*domain.Rack, errors.IBaseError)
	FindByCode(code string) (*domain.Rack, errors.IBaseError)
	Create(instance *domain.Rack) (*domain.Rack, errors.IBaseError)
	Update(instance *domain.Rack) (*domain.Rack, errors.IBaseError)
	Delete(instance *domain.Rack) errors.IBaseError
}
