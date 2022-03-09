package application

import (
	"warehouse/domain"
	"warehouse/domain/repository"
	"warehouse/infraestructure/errors"
)

type MarketApplication struct {
	marketRepository repository.IMarketRepository
}

func NewMarketApplication(marketRepository repository.IMarketRepository) *MarketApplication {
	return &MarketApplication{ marketRepository }
}

func (marketApplication *MarketApplication) All() (*[]domain.Market, errors.IBaseError) {
	return marketApplication.marketRepository.All()
}

func (marketApplication *MarketApplication) Show(id uint) (*domain.Market, errors.IBaseError) {
	return marketApplication.marketRepository.Find(id)
}

func (marketApplication *MarketApplication) Update(market *domain.Market) (*domain.Market, errors.IBaseError) {
	return marketApplication.marketRepository.Update(market)
}

func (marketApplication *MarketApplication) Create(market *domain.Market) (*domain.Market, errors.IBaseError) {
	return marketApplication.marketRepository.Create(market)
}

func (marketApplication *MarketApplication) Delete(market *domain.Market) errors.IBaseError {
	return marketApplication.marketRepository.Delete(market)
}