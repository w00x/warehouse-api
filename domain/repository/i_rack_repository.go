package repository

import (
	"warehouse/domain"
	"warehouse/infraestructure/errors"
)

type IRackRepository interface {
	All() 							(*[]domain.Rack, errors.IBaseError)
	Find(id uint) 					(*domain.Rack, errors.IBaseError)
	Create(instance *domain.Rack) 	(*domain.Rack, errors.IBaseError)
	Update(instance *domain.Rack) 	(*domain.Rack, errors.IBaseError)
	Delete(instance *domain.Rack) 	errors.IBaseError
}
