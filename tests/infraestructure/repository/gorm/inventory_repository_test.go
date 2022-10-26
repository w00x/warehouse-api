package gorm

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"warehouse/infrastructure/repository/gorm"
	"warehouse/shared"
	"warehouse/tests/factories"
)

func TestInventoryRepository_All(t *testing.T) {
	sizeOfInventories := 5
	inventories := factories.NewInventoryFactoryList(sizeOfInventories)
	inventoryRepo := gorm.NewInventoryRepository()
	allInventories, err := inventoryRepo.All()
	assert.Nil(t, err)

	var inventoriesIds []string
	for _, inventory := range *allInventories {
		inventoriesIds = append(inventoriesIds, inventory.Id())
	}

	assert.Contains(t, inventoriesIds, inventories[0].Id())
	assert.Contains(t, inventoriesIds, inventories[1].Id())
	assert.Contains(t, inventoriesIds, inventories[2].Id())
}

func TestInventoryRepository_Create(t *testing.T) {
	inventoryRepo := gorm.NewInventoryRepository()
	inventoryData := factories.NewInventoryDomainFactory()
	dateTime := inventoryData.OperationDate
	inventory, err := inventoryRepo.Create(inventoryData)
	assert.Nil(t, err)

	assert.Equal(t, inventory.OperationDate, dateTime)
	assert.NotNil(t, inventory.Id())
}

func TestInventoryRepository_Delete(t *testing.T) {
	inventory := factories.NewInventoryFactory()
	inventoryId := inventory.Id()
	assert.NotNil(t, inventoryId)

	inventoryRepo := gorm.NewInventoryRepository()
	inventoryFounded, err := inventoryRepo.Find(inventoryId)
	assert.Nil(t, err)
	assert.NotNil(t, inventoryFounded)
	assert.Equal(t, inventoryId, inventoryFounded.Id())

	assert.Nil(t, inventoryRepo.Delete(inventory))
	inventoryFounded, err = inventoryRepo.Find(inventoryId)
	assert.NotNil(t, err)
	assert.Nil(t, inventoryFounded)
}

func TestInventoryRepository_Find(t *testing.T) {
	inventory := factories.NewInventoryFactory()
	inventoryId := inventory.Id()
	assert.NotNil(t, inventoryId)

	inventoryRepo := gorm.NewInventoryRepository()
	inventoryFounded, err := inventoryRepo.Find(inventoryId)
	assert.Nil(t, err)
	assert.NotNil(t, inventoryFounded)
	assert.Equal(t, inventoryId, inventoryFounded.Id())
}

func TestInventoryRepository_Update(t *testing.T) {
	inventory := factories.NewInventoryFactory()
	newDate := shared.DateTime(time.Time{}.Add(-time.Hour * 24))
	inventoryRepo := gorm.NewInventoryRepository()
	inventory.OperationDate = newDate
	inventoryUpdated, err := inventoryRepo.Update(inventory)

	assert.Nil(t, err)
	assert.NotNil(t, inventoryUpdated)

	inventoryFounded, err := inventoryRepo.Find(inventoryUpdated.Id())
	assert.Nil(t, err)
	assert.NotNil(t, inventoryFounded)
	assert.Equal(t, newDate, inventoryFounded.OperationDate)
}

func TestNewInventoryRepository(t *testing.T) {
	repo := gorm.NewInventoryRepository()
	assert.NotNil(t, repo)
	assert.IsType(t, gorm.InventoryRepository{}, *repo)
}
