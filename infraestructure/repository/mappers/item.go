package mappers

import (
	"warehouse/domain"
	"warehouse/infraestructure/repository/models"
)

func FromItemDomainToModel(i *domain.Item) *models.Item {
	if i == nil {
		return nil
	}
	return models.NewItem(i.Id, i.Name, i.UnitSizePresentation, i.SizePresentation, i.Code, i.Container, i.Photo)
}

func FromItemModelToDomain(i *models.Item) *domain.Item {
	return domain.NewItem(i.Id, i.Name, i.UnitSizePresentation, i.SizePresentation, i.Code, i.Container, i.Photo)
}