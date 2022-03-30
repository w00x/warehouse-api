package mappers

import (
	"warehouse/domain"
	"warehouse/infraestructure/repository/models"
)

func FromMarketDomainToModel(i *domain.Market) *models.Market {
	if i == nil {
		return nil
	}
	return models.NewMarket(i.Id, i.Name)
}

func FromMarketModelToDomain(i *models.Market) *domain.Market {
	return domain.NewMarket(i.Id, i.Name)
}