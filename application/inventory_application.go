package application

import (
	"warehouse/domain"
	"warehouse/domain/repository"
	"warehouse/infrastructure/errors"
)

type InventoryApplication struct {
	inventoryRepository repository.IInventoryRepository
}

func NewInventoryApplication(inventoryRepository repository.IInventoryRepository) *InventoryApplication {
	return &InventoryApplication{inventoryRepository}
}

func (inventoryApplication *InventoryApplication) All() (*[]domain.Inventory, errors.IBaseError) {
	return inventoryApplication.inventoryRepository.All()
}

func (inventoryApplication *InventoryApplication) Show(id string) (*domain.Inventory, errors.IBaseError) {
	return inventoryApplication.inventoryRepository.Find(id)
}

func (inventoryApplication *InventoryApplication) Update(inventory *domain.Inventory) (*domain.Inventory, errors.IBaseError) {
	return inventoryApplication.inventoryRepository.Update(inventory)
}

func (inventoryApplication *InventoryApplication) Create(inventory *domain.Inventory) (*domain.Inventory, errors.IBaseError) {
	return inventoryApplication.inventoryRepository.Create(inventory)
}

func (inventoryApplication *InventoryApplication) Delete(inventory *domain.Inventory) errors.IBaseError {
	return inventoryApplication.inventoryRepository.Delete(inventory)
}
