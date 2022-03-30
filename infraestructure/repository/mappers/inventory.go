package mappers

import (
	"warehouse/domain"
	"warehouse/infraestructure/repository/models"
)

func FromInventoryDomainToModel(i *domain.Inventory) *models.Inventory {
	return models.NewInventory(i.Id, i.OperationDate)
}

func FromInventoryModelToDomain(i *models.Inventory) *domain.Inventory {
	return domain.NewInventory(i.Id, i.OperationDate)
}