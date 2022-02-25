package repository

import (
	"time"
	"warehouse/domain"
)

type IPriceRepository interface {
	All() ([]*domain.Price, error)
	Find(id string) (*domain.Price, error)
	Update(id string, market *domain.Market, item *domain.Item, price int,
		date time.Time) error
	Create(market *domain.Market, item *domain.Item, price int,
		date time.Time) (*domain.Price, error)
	Delete(id string) error
}
