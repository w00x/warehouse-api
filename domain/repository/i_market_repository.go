package repository

import "warehouse/domain"

type IMarketRepository interface {
	All() ([]*domain.Market, error)
	Find(id string) (*domain.Market, error)
	Update(id string, name string) error
	Create(name string) (*domain.Market, error)
	Delete(id string) error
}
