package application

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"warehouse/application"
	"warehouse/infrastructure/errors"
	"warehouse/infrastructure/repository/gorm"
	"warehouse/tests/factories"
)

func TestPriceApplication_All(t *testing.T) {
	sizeOfInventories := 5
	pricesList := factories.NewPriceFactoryList(sizeOfInventories)
	repo := gorm.NewPriceRepository()
	priceApplication := application.NewPriceApplication(repo)
	prices, err := priceApplication.All()

	assert.Nil(t, err)
	var ids []string
	for _, price := range pricesList {
		ids = append(ids, price.Id())
	}

	assert.Contains(t, ids, (*prices)[0].Id())
}

func TestPriceApplication_Create(t *testing.T) {
	price := factories.NewPriceDomainFactory()
	repo := gorm.NewPriceRepository()
	priceApplication := application.NewPriceApplication(repo)
	newPrice, err := priceApplication.Create(price)

	assert.Nil(t, err)
	assert.NotNil(t, newPrice.Id())
	assert.Equal(t, price.Price, newPrice.Price)
}

func TestPriceApplication_Delete(t *testing.T) {
	price := factories.NewPriceFactory()
	repo := gorm.NewPriceRepository()
	priceApplication := application.NewPriceApplication(repo)
	err := priceApplication.Delete(price)
	assert.Nil(t, err)

	findedPrice, errFind := repo.Find(price.Id())

	assert.NotNil(t, errFind)
	assert.IsType(t, errFind, errors.NewNotFoundError(""))
	assert.Nil(t, findedPrice)
}

func TestPriceApplication_Show(t *testing.T) {
	price := factories.NewPriceFactory()
	repo := gorm.NewPriceRepository()
	priceApplication := application.NewPriceApplication(repo)
	findedPrice, err := priceApplication.Show(price.Id())

	assert.Nil(t, err)
	assert.Equal(t, price.Id(), findedPrice.Id())
}

func TestPriceApplication_Update(t *testing.T) {
	price := factories.NewPriceFactory()
	values := factories.NewPriceObjectFactory()

	priceValue := values["price"].(float64)
	repo := gorm.NewPriceRepository()
	priceApplication := application.NewPriceApplication(repo)
	price.Price = priceValue
	updatedPrice, errUpdate := priceApplication.Update(price)

	assert.Nil(t, errUpdate)
	assert.Equal(t, updatedPrice.Price, priceValue)
}

func TestNewPriceApplication(t *testing.T) {
	repo := gorm.NewPriceRepository()
	priceApplication := application.NewPriceApplication(repo)

	assert.NotNil(t, priceApplication)
}
