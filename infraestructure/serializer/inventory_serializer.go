package serializer

import (
	"time"
	"warehouse/domain"
	"warehouse/shared"
)

type InventorySerializer struct {
	Id 				string `json:"id"`
	OperationDate 	shared.DateTime `json:"operation_date" binding:"required"`
}

func NewInventorySerializer(id string, operationDate shared.DateTime) *InventorySerializer {
	return &InventorySerializer{Id: id, OperationDate: operationDate}
}

func NewInventorySerializerFromDomain(inventory *domain.Inventory) *InventorySerializer {
	return NewInventorySerializer(inventory.Id, shared.DateTime(inventory.OperationDate))
}

func NewInventoryListSerializerFromDomains(inventories []*domain.Inventory) []*InventorySerializer {
	var inventorySerializers []*InventorySerializer
	for _, inventory := range inventories {
		inventorySerializers = append(inventorySerializers,
			NewInventorySerializerFromDomain(inventory))
	}
	return inventorySerializers
}

func (ise InventorySerializer) ToDomain() *domain.Inventory {
	return domain.NewInventory(ise.Id, time.Time(ise.OperationDate))
}
