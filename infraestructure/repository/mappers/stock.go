package mappers

import (
	"warehouse/domain"
	"warehouse/infraestructure/repository/models"
)

func FromStockDomainToModel(i *domain.Stock) *models.Stock {
	return models.NewStock(i.Id, FromItemDomainToModel(i.Item), FromRackDomainToModel(i.Rack), i.Quantity, i.OperationDate, i.ExpirationDate, i.ItemId, i.RackId)
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
	return domain.NewStock(i.Id, item, rack, i.Quantity, i.OperationDate, i.ExpirationDate, i.RackId, i.ItemId)
}