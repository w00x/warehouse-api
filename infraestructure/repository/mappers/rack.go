package mappers

import (
	"warehouse/domain"
	"warehouse/infraestructure/repository/models"
)

func FromRackDomainToModel(i *domain.Rack) *models.Rack {
	if i == nil {
		return nil
	}
	return models.NewRack(i.Id, i.Name, i.Code)
}

func FromRackModelToDomain(i *models.Rack) *domain.Rack {
	return domain.NewRack(i.Id, i.Name, i.Code)
}