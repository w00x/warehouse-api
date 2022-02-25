package application

import (
	"time"
	"warehouse/domain"
	"warehouse/domain/repository"
)

type PriceApplication struct {
	priceRepository repository.IPriceRepository
}

func NewPriceApplication(priceRepository repository.IPriceRepository) *PriceApplication {
	return &PriceApplication{ priceRepository }
}

func (priceApplication *PriceApplication) All() ([]*domain.Price, error) {
	return priceApplication.priceRepository.All()
}

func (priceApplication *PriceApplication) Show(id string) (*domain.Price, error) {
	return priceApplication.priceRepository.Find(id)
}

func (priceApplication *PriceApplication) Update(id string, market *domain.Market,
	item *domain.Item, price int, date time.Time) error {
	return priceApplication.priceRepository.Update(id, market, item, price, date)
}

func (priceApplication *PriceApplication) Create(market *domain.Market,
	item *domain.Item, price int, date time.Time) (*domain.Price, error) {
	return priceApplication.priceRepository.Create(market, item, price, date)
}

func (priceApplication *PriceApplication) Delete(id string) error {
	return priceApplication.priceRepository.Delete(id)
}