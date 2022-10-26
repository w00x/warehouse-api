package application

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"warehouse/application"
	"warehouse/infrastructure/errors"
	"warehouse/infrastructure/repository/gorm"
	"warehouse/tests/factories"
)

func TestItemApplication_All(t *testing.T) {
	sizeOfInventories := 5
	itemsList := factories.NewItemFactoryList(sizeOfInventories)
	repo := gorm.NewItemRepository()
	itemApplication := application.NewItemApplication(repo)
	items, err := itemApplication.All()

	assert.Nil(t, err)
	var ids []string
	for _, item := range itemsList {
		ids = append(ids, item.Id())
	}

	assert.Contains(t, ids, (*items)[0].Id())
}

func TestItemApplication_Create(t *testing.T) {
	item := factories.NewItemDomainFactory()
	repo := gorm.NewItemRepository()
	itemApplication := application.NewItemApplication(repo)
	newItem, err := itemApplication.Create(item)

	assert.Nil(t, err)
	assert.NotNil(t, newItem.Id())
	assert.Equal(t, item.Code, newItem.Code)
}

func TestItemApplication_Delete(t *testing.T) {
	item := factories.NewItemFactory()
	repo := gorm.NewItemRepository()
	itemApplication := application.NewItemApplication(repo)
	err := itemApplication.Delete(item)
	assert.Nil(t, err)

	findedItem, errFind := repo.Find(item.Id())

	assert.NotNil(t, errFind)
	assert.IsType(t, errFind, errors.NewNotFoundError(""))
	assert.Nil(t, findedItem)
}

func TestItemApplication_Show(t *testing.T) {
	item := factories.NewItemFactory()
	repo := gorm.NewItemRepository()
	itemApplication := application.NewItemApplication(repo)
	findedItem, err := itemApplication.Show(item.Id())

	assert.Nil(t, err)
	assert.Equal(t, item.Id(), findedItem.Id())
}

func TestItemApplication_Update(t *testing.T) {
	item := factories.NewItemFactory()
	values := factories.NewItemObjectFactory()

	code := values["code"].(string)
	repo := gorm.NewItemRepository()
	itemApplication := application.NewItemApplication(repo)
	item.Code = code
	updatedItem, errUpdate := itemApplication.Update(item)

	assert.Nil(t, errUpdate)
	assert.Equal(t, updatedItem.Code, code)
}

func TestNewItemApplication(t *testing.T) {
	repo := gorm.NewItemRepository()
	itemApplication := application.NewItemApplication(repo)

	assert.NotNil(t, itemApplication)
}
