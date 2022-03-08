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

func (inventoryApplication *InventoryApplication) All() (*[]domain.Inventory, errors.IBaseError) {
	return inventoryApplication.inventoryRepository.All()
}

func (inventoryApplication *InventoryApplication) Show(id uint) (*domain.Inventory, errors.IBaseError) {
	return inventoryApplication.inventoryRepository.Find(id)
}

func (inventoryApplication *InventoryApplication) Update(id uint, operationDate time.Time) (*domain.Inventory, errors.IBaseError) {
	inventory, err := inventoryApplication.inventoryRepository.Find(id)
	if err != nil {
		return nil, err
	}
	inventory.OperationDate = operationDate
	return inventoryApplication.inventoryRepository.Update(inventory)
}

func (inventoryApplication *InventoryApplication) Create(operationDate time.Time) (*domain.Inventory, errors.IBaseError) {
	inventory := domain.NewInventory(0, operationDate)
	return inventoryApplication.inventoryRepository.Create(inventory)
}

func (inventoryApplication *InventoryApplication) Delete(id uint) errors.IBaseError {
	inventory, err := inventoryApplication.inventoryRepository.Find(id)
	if err != nil {
		return err
	}
	return inventoryApplication.inventoryRepository.Delete(inventory)
}