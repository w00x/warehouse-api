package mappers

import (
	"warehouse/domain"
	"warehouse/infrastructure/repository/models"
)

func FromItemDomainToModel(i *domain.Item) *models.Item {
	if i == nil {
		return nil
	}
	return models.NewItem(i.Id(), i.Name, i.UnitSizePresentation, i.SizePresentation, i.Code, i.Container, i.Photo)
}

func FromItemModelToDomain(i *models.Item) *domain.Item {
	return domain.NewItem(i.ID, i.Name, i.UnitSizePresentation, i.SizePresentation, i.Code, i.Container, i.Photo)
}

func NewItemListDomainFromModel(items *[]models.Item) *[]domain.Item {
	var itemsDomain []domain.Item
	for _, item := range *items {
		itemsDomain = append(itemsDomain, *FromItemModelToDomain(&item))
	}
	return &itemsDomain
}
