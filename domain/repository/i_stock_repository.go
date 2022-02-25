package repository

import (
	"time"
	"warehouse/domain"
)

type IStockRepository interface {
	All() ([]*domain.Stock, error)
	Find(id string) (*domain.Stock, error)
	Update(id string, item *domain.Item, rack *domain.Rack, quantity int, operationDate time.Time,
		expirationDate time.Time) error
	Create(item *domain.Item, rack *domain.Rack, quantity int, operationDate time.Time,
		expirationDate time.Time) (*domain.Stock, error)
	Delete(id string) error
}
