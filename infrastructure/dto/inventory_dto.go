package dto

import (
	"warehouse/domain"
	"warehouse/shared"
)

type InventoryDto struct {
	Id            string          `json:"id" uri:"id"`
	OperationDate shared.DateTime `json:"operation_date" binding:"required"`
}

func NewInventoryDto(id string, operationDate shared.DateTime) *InventoryDto {
	return &InventoryDto{Id: id, OperationDate: operationDate}
}

func NewInventoryDtoFromDomain(inventory *domain.Inventory) *InventoryDto {
	return NewInventoryDto(inventory.Id(), inventory.OperationDate)
}

func NewInventoryListDtoFromDomains(inventories *[]domain.Inventory) []*InventoryDto {
	var inventoryDtos []*InventoryDto
	for _, inventory := range *inventories {
		inventoryDtos = append(inventoryDtos,
			NewInventoryDtoFromDomain(&inventory))
	}
	return inventoryDtos
}

func (ise InventoryDto) ToDomain() *domain.Inventory {
	return domain.NewInventory(ise.Id, ise.OperationDate)
}
