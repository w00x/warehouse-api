package application

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"warehouse/application"
	"warehouse/infrastructure/errors"
	"warehouse/infrastructure/repository/gorm"
	"warehouse/shared"
	"warehouse/tests/factories"
)

func TestInventoryApplication_All(t *testing.T) {
	sizeOfInventories := 5
	inventoriesList := factories.NewInventoryFactoryList(sizeOfInventories)
	repo := gorm.NewInventoryRepository()
	inventoryApplication := application.NewInventoryApplication(repo)
	inventories, err := inventoryApplication.All()

	assert.Nil(t, err)
	var ids []string
	for _, inventory := range inventoriesList {
		ids = append(ids, inventory.Id())
	}

	assert.Contains(t, ids, (*inventories)[0].Id())
}

func TestInventoryApplication_Create(t *testing.T) {
	inventory := factories.NewInventoryDomainFactory()
	repo := gorm.NewInventoryRepository()
	inventoryApplication := application.NewInventoryApplication(repo)
	newInventory, err := inventoryApplication.Create(inventory)

	assert.Nil(t, err)
	assert.NotNil(t, newInventory.Id())
	assert.Equal(t, inventory.OperationDate, newInventory.OperationDate)
}

func TestInventoryApplication_Delete(t *testing.T) {
	inventory := factories.NewInventoryFactory()
	repo := gorm.NewInventoryRepository()
	inventoryApplication := application.NewInventoryApplication(repo)
	err := inventoryApplication.Delete(inventory)
	assert.Nil(t, err)

	findedInventory, errFind := repo.Find(inventory.Id())

	assert.NotNil(t, errFind)
	assert.IsType(t, errFind, errors.NewNotFoundError(""))
	assert.Nil(t, findedInventory)
}

func TestInventoryApplication_Show(t *testing.T) {
	inventory := factories.NewInventoryFactory()
	repo := gorm.NewInventoryRepository()
	inventoryApplication := application.NewInventoryApplication(repo)
	findedInventory, err := inventoryApplication.Show(inventory.Id())

	assert.Nil(t, err)
	assert.Equal(t, inventory.Id(), findedInventory.Id())
}

func TestInventoryApplication_Update(t *testing.T) {
	inventory := factories.NewInventoryFactory()
	values := factories.NewInventoryObjectFactory()

	operationDate, err := shared.StringToDate(values["operation_date"].(string))
	assert.Nil(t, err)
	repo := gorm.NewInventoryRepository()
	inventoryApplication := application.NewInventoryApplication(repo)
	inventory.OperationDate = *operationDate
	updatedInventory, errUpdate := inventoryApplication.Update(inventory)

	assert.Nil(t, errUpdate)
	assert.Equal(t, updatedInventory.OperationDate, *operationDate)
}

func TestNewInventoryApplication(t *testing.T) {
	repo := gorm.NewInventoryRepository()
	inventoryApplication := application.NewInventoryApplication(repo)

	assert.NotNil(t, inventoryApplication)
}
