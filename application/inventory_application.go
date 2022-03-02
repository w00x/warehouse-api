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

func (inventoryApplication *InventoryApplication) All() ([]*domain.Inventory, errors.IBaseError) {
	return inventoryApplication.inventoryRepository.All()
}

func (inventoryApplication *InventoryApplication) Show(id string) (*domain.Inventory, errors.IBaseError) {
	return inventoryApplication.inventoryRepository.Find(id)
}

func (inventoryApplication *InventoryApplication) Update(id string, operationDate time.Time) errors.IBaseError {
	return inventoryApplication.inventoryRepository.Update(id, operationDate)
}

func (inventoryApplication *InventoryApplication) Create(operationDate time.Time) (*domain.Inventory, errors.IBaseError) {
	return inventoryApplication.inventoryRepository.Create(operationDate)
}

func (inventoryApplication *InventoryApplication) Delete(id string) errors.IBaseError {
	return inventoryApplication.inventoryRepository.Delete(id)
}