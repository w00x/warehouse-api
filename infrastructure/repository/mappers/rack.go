package mappers

import (
	"warehouse/domain"
	"warehouse/infrastructure/repository/models"
)

func FromRackDomainToModel(i *domain.Rack) *models.Rack {
	if i == nil {
		return nil
	}
	return models.NewRack(i.Id(), i.Name, i.Code)
}

func FromRackModelToDomain(i *models.Rack) *domain.Rack {
	return domain.NewRack(i.ID, i.Name, i.Code)
}

func NewRackListDomainFromModel(racks *[]models.Rack) *[]domain.Rack {
	var racksDomain []domain.Rack
	for _, rack := range *racks {
		racksDomain = append(racksDomain, *FromRackModelToDomain(&rack))
	}
	return &racksDomain
}
