package mappers

import (
	"warehouse/domain"
	"warehouse/infrastructure/repository/models"
)

func FromStockDomainToModel(i *domain.Stock) *models.Stock {
	return models.NewStock(i.Id(), FromItemDomainToModel(i.Item), FromRackDomainToModel(i.Rack),
		i.Quantity, i.OperationDate, i.ExpirationDate, i.ItemId, i.RackId, i.Comment)
}

func FromStockModelToDomain(i *models.Stock) *domain.Stock {
	var rack *domain.Rack
	var item *domain.Item
	if i.Rack != nil {
		rack = FromRackModelToDomain(i.Rack)
	}
	if i.Item != nil {
		item = FromItemModelToDomain(i.Item)
	}
	return domain.NewStock(i.ID, item, rack, i.Quantity, i.OperationDate, i.ExpirationDate,
		i.Comment, i.ItemId, i.RackId)
}

func NewStockListDomainFromModel(stocks *[]models.Stock) *[]domain.Stock {
	var stocksDomain []domain.Stock
	for _, stock := range *stocks {
		stocksDomain = append(stocksDomain, *FromStockModelToDomain(&stock))
	}
	return &stocksDomain
}
