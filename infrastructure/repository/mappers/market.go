package mappers

import (
	"warehouse/domain"
	"warehouse/infrastructure/repository/models"
)

func FromMarketDomainToModel(i *domain.Market) *models.Market {
	if i == nil {
		return nil
	}
	return models.NewMarket(i.Id(), i.Name)
}

func FromMarketModelToDomain(i *models.Market) *domain.Market {
	return domain.NewMarket(i.ID, i.Name)
}

func NewMarketListDomainFromModel(markets *[]models.Market) *[]domain.Market {
	var marketsDomain []domain.Market
	for _, market := range *markets {
		marketsDomain = append(marketsDomain, *FromMarketModelToDomain(&market))
	}
	return &marketsDomain
}
