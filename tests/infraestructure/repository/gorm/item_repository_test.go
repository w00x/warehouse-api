package gorm

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"testing"
	"warehouse/infrastructure/repository/gorm"
	"warehouse/tests/factories"
)

func TestItemRepository_All(t *testing.T) {
	sizeOfItems := 5
	items := factories.NewItemFactoryList(sizeOfItems)
	itemRepo := gorm.NewItemRepository()
	allItems, err := itemRepo.All()
	assert.Nil(t, err)

	var itemsIds []string
	for _, item := range *allItems {
		itemsIds = append(itemsIds, item.Id())
	}

	assert.Contains(t, itemsIds, items[0].Id())
	assert.Contains(t, itemsIds, items[1].Id())
	assert.Contains(t, itemsIds, items[2].Id())
}

func TestItemRepository_Create(t *testing.T) {
	itemRepo := gorm.NewItemRepository()
	itemData := factories.NewItemDomainFactory()

	item, err := itemRepo.Create(itemData)
	assert.Nil(t, err)

	assert.Equal(t, item.Name, itemData.Name)
	assert.Equal(t, item.UnitSizePresentation, itemData.UnitSizePresentation)
	assert.Equal(t, item.SizePresentation, itemData.SizePresentation)
	assert.NotNil(t, item.Id())
}

func TestItemRepository_Delete(t *testing.T) {
	item := factories.NewItemFactory()
	itemId := item.Id()
	assert.NotNil(t, itemId)

	itemRepo := gorm.NewItemRepository()
	itemFounded, err := itemRepo.Find(itemId)
	assert.Nil(t, err)
	assert.NotNil(t, itemFounded)
	assert.Equal(t, itemId, itemFounded.Id())

	assert.Nil(t, itemRepo.Delete(item))
	itemFounded, err = itemRepo.Find(itemId)
	assert.NotNil(t, err)
	assert.Nil(t, itemFounded)
}

func TestItemRepository_Find(t *testing.T) {
	item := factories.NewItemFactory()
	itemId := item.Id()
	assert.NotNil(t, itemId)

	itemRepo := gorm.NewItemRepository()
	itemFounded, err := itemRepo.Find(itemId)
	assert.Nil(t, err)
	assert.NotNil(t, itemFounded)
	assert.Equal(t, itemId, itemFounded.Id())
}

func TestItemRepository_Update(t *testing.T) {
	item := factories.NewItemFactory()
	newName := gofakeit.LoremIpsumWord()
	itemRepo := gorm.NewItemRepository()
	item.Name = newName
	itemUpdated, err := itemRepo.Update(item)

	assert.Nil(t, err)
	assert.NotNil(t, itemUpdated)

	itemFounded, err := itemRepo.Find(itemUpdated.Id())
	assert.Nil(t, err)
	assert.NotNil(t, itemFounded)
	assert.Equal(t, newName, itemFounded.Name)
}

func TestNewItemRepository(t *testing.T) {
	repo := gorm.NewItemRepository()
	assert.NotNil(t, repo)
	assert.IsType(t, gorm.ItemRepository{}, *repo)
}
