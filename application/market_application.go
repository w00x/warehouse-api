package application

import (
	"warehouse/domain"
	"warehouse/domain/repository"
)

type MarketApplication struct {
	marketRepository repository.IMarketRepository
}

func NewMarketApplication(marketRepository repository.IMarketRepository) *MarketApplication {
	return &MarketApplication{ marketRepository }
}

func (marketApplication *MarketApplication) All() ([]*domain.Market, error) {
	return marketApplication.marketRepository.All()
}

func (marketApplication *MarketApplication) Show(id string) (*domain.Market, error) {
	return marketApplication.marketRepository.Find(id)
}

func (marketApplication *MarketApplication) Update(id string, name string) error {
	return marketApplication.marketRepository.Update(id, name)
}

func (marketApplication *MarketApplication) Create(name string) (*domain.Market, error) {
	return marketApplication.marketRepository.Create(name)
}

func (marketApplication *MarketApplication) Delete(id string) error {
	return marketApplication.marketRepository.Delete(id)
}