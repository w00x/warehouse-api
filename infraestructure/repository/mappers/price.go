package mappers

import (
	"warehouse/domain"
	"warehouse/infraestructure/repository/models"
)

func FromPriceDomainToModel(i *domain.Price) *models.Price {
	return models.NewPrice(i.Id, FromMarketDomainToModel(i.Market), FromItemDomainToModel(i.Item), i.Price, i.Date, i.ItemId, i.MarketId)
}

func FromPriceModelToDomain(i *models.Price) *domain.Price {
	var market *domain.Market
	var item *domain.Item
	if i.Market != nil {
		market = FromMarketModelToDomain(i.Market)
	}
	if i.Item != nil {
		item = FromItemModelToDomain(i.Item)
	}
	return domain.NewPrice(i.Id, market, item, i.Price, i.Date, i.ItemId, i.MarketId)
}