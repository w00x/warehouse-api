package mappers

import (
	"warehouse/domain"
	"warehouse/infrastructure/repository/models"
)

func FromInventoryDomainToModel(i *domain.Inventory) *models.Inventory {
	return models.NewInventory(i.Id(), i.OperationDate)
}

func FromInventoryModelToDomain(i *models.Inventory) *domain.Inventory {
	return domain.NewInventory(i.ID, i.OperationDate)
}

func NewInventoryListDomainFromModel(inventories *[]models.Inventory) *[]domain.Inventory {
	var inventoriesDomain []domain.Inventory
	for _, inventory := range *inventories {
		inventoriesDomain = append(inventoriesDomain, *FromInventoryModelToDomain(&inventory))
	}
	return &inventoriesDomain
}
