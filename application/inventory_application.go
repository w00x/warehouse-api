package application

import (
	"time"
	"warehouse/domain"
	"warehouse/domain/repository"
	"warehouse/infraestructure/errors"
)

type InventoryApplication struct {
	inventoryRepository repository.IInventoryRepository
}

func NewInventoryApplication(inventoryRepository repository.IInventoryRepository) *InventoryApplication {
	return &InventoryApplication{ inventoryRepository }
}

func (inventoryApplication *InventoryApplication) All() ([]*domain.Inventory, error) {
	return inventoryApplication.inventoryRepository.All()
}

func (inventoryApplication *InventoryApplication) Show(id string) (*domain.Inventory, error) {
	return inventoryApplication.inventoryRepository.Find(id)
}

func (inventoryApplication *InventoryApplication) Update(id string, operationDate time.Time) errors.BaseError {
	return inventoryApplication.inventoryRepository.Update(id, operationDate)
}

func (inventoryApplication *InventoryApplication) Create(operationDate time.Time) (*domain.Inventory, error) {
	return inventoryApplication.inventoryRepository.Create(operationDate)
}

func (inventoryApplication *InventoryApplication) Delete(id string) error {
	return inventoryApplication.inventoryRepository.Delete(id)
}