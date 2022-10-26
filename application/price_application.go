package application

import (
	"warehouse/domain"
	"warehouse/domain/repository"
	"warehouse/infrastructure/errors"
)

type PriceApplication struct {
	priceRepository repository.IPriceRepository
}

func NewPriceApplication(priceRepository repository.IPriceRepository) *PriceApplication {
	return &PriceApplication{priceRepository}
}

func (priceApplication *PriceApplication) All() (*[]domain.Price, errors.IBaseError) {
	return priceApplication.priceRepository.All()
}

func (priceApplication *PriceApplication) Show(id string) (*domain.Price, errors.IBaseError) {
	return priceApplication.priceRepository.Find(id)
}

func (priceApplication *PriceApplication) Update(price *domain.Price) (*domain.Price, errors.IBaseError) {
	return priceApplication.priceRepository.Update(price)
}

func (priceApplication *PriceApplication) Create(price *domain.Price) (*domain.Price, errors.IBaseError) {
	return priceApplication.priceRepository.Create(price)
}

func (priceApplication *PriceApplication) Delete(price *domain.Price) errors.IBaseError {
	return priceApplication.priceRepository.Delete(price)
}
