package serializer

import (
	"warehouse/domain"
	"warehouse/shared"
)

type InventorySerializer struct {
	Id 				uint `json:"id" uri:"id"`
	OperationDate 	shared.DateTime `json:"operation_date" binding:"required"`
}

func NewInventorySerializer(id uint, operationDate shared.DateTime) *InventorySerializer {
	return &InventorySerializer{Id: id, OperationDate: operationDate}
}

func NewInventorySerializerFromDomain(inventory *domain.Inventory) *InventorySerializer {
	return NewInventorySerializer(inventory.Id, shared.DateTime(inventory.OperationDate))
}

func NewInventoryListSerializerFromDomains(inventories *[]domain.Inventory) []*InventorySerializer {
	var inventorySerializers []*InventorySerializer
	for _, inventory := range *inventories {
		inventorySerializers = append(inventorySerializers,
			NewInventorySerializerFromDomain(&inventory))
	}
	return inventorySerializers
}

func (ise InventorySerializer) ToDomain() *domain.Inventory {
	return domain.NewInventory(ise.Id, ise.OperationDate)
}
