package repository

import (
	"warehouse/domain"
	"warehouse/infraestructure/errors"
)

type IStockRepository interface {
	All() 							(*[]domain.Stock, errors.IBaseError)
	Find(id uint) 					(*domain.Stock, errors.IBaseError)
	Create(instance *domain.Stock) 	(*domain.Stock, errors.IBaseError)
	Update(instance *domain.Stock) 	(*domain.Stock, errors.IBaseError)
	Delete(instance *domain.Stock) 	errors.IBaseError
}
