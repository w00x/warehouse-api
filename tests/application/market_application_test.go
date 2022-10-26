package application

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"warehouse/application"
	"warehouse/infrastructure/errors"
	"warehouse/infrastructure/repository/gorm"
	"warehouse/tests/factories"
)

func TestMarketApplication_All(t *testing.T) {
	sizeOfInventories := 5
	martketsList := factories.NewMarketFactoryList(sizeOfInventories)
	repo := gorm.NewMarketRepository()
	martketApplication := application.NewMarketApplication(repo)
	martkets, err := martketApplication.All()

	assert.Nil(t, err)
	var ids []string
	for _, martket := range martketsList {
		ids = append(ids, martket.Id())
	}

	assert.Contains(t, ids, (*martkets)[0].Id())
}

func TestMarketApplication_Create(t *testing.T) {
	martket := factories.NewMarketDomainFactory()
	repo := gorm.NewMarketRepository()
	martketApplication := application.NewMarketApplication(repo)
	newMarket, err := martketApplication.Create(martket)

	assert.Nil(t, err)
	assert.NotNil(t, newMarket.Id())
	assert.Equal(t, martket.Name, newMarket.Name)
}

func TestMarketApplication_Delete(t *testing.T) {
	martket := factories.NewMarketFactory()
	repo := gorm.NewMarketRepository()
	martketApplication := application.NewMarketApplication(repo)
	err := martketApplication.Delete(martket)
	assert.Nil(t, err)

	findedMarket, errFind := repo.Find(martket.Id())

	assert.NotNil(t, errFind)
	assert.IsType(t, errFind, errors.NewNotFoundError(""))
	assert.Nil(t, findedMarket)
}

func TestMarketApplication_Show(t *testing.T) {
	martket := factories.NewMarketFactory()
	repo := gorm.NewMarketRepository()
	martketApplication := application.NewMarketApplication(repo)
	findedMarket, err := martketApplication.Show(martket.Id())

	assert.Nil(t, err)
	assert.Equal(t, martket.Id(), findedMarket.Id())
}

func TestMarketApplication_Update(t *testing.T) {
	martket := factories.NewMarketFactory()
	values := factories.NewMarketObjectFactory()

	name := values["name"].(string)
	repo := gorm.NewMarketRepository()
	martketApplication := application.NewMarketApplication(repo)
	martket.Name = name
	updatedMarket, errUpdate := martketApplication.Update(martket)

	assert.Nil(t, errUpdate)
	assert.Equal(t, updatedMarket.Name, name)
}

func TestNewMarketApplication(t *testing.T) {
	repo := gorm.NewMarketRepository()
	martketApplication := application.NewMarketApplication(repo)

	assert.NotNil(t, martketApplication)
}
