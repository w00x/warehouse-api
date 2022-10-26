package gorm

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"testing"
	"warehouse/infrastructure/repository/gorm"
	"warehouse/tests/factories"
)

func TestPriceRepository_All(t *testing.T) {
	sizeOfPrices := 5
	prices := factories.NewPriceFactoryList(sizeOfPrices)
	priceRepo := gorm.NewPriceRepository()
	allPrices, err := priceRepo.All()
	assert.Nil(t, err)

	var pricesIds []string
	for _, price := range *allPrices {
		pricesIds = append(pricesIds, price.Id())
	}

	assert.Contains(t, pricesIds, prices[0].Id())
	assert.Contains(t, pricesIds, prices[1].Id())
	assert.Contains(t, pricesIds, prices[2].Id())
}

func TestPriceRepository_Create(t *testing.T) {
	priceRepo := gorm.NewPriceRepository()
	priceData := factories.NewPriceDomainFactory()

	price, err := priceRepo.Create(priceData)
	assert.Nil(t, err)

	assert.Equal(t, price.Price, priceData.Price)
	assert.NotNil(t, price.Id())
}

func TestPriceRepository_Delete(t *testing.T) {
	price := factories.NewPriceFactory()
	priceId := price.Id()
	assert.NotNil(t, priceId)

	priceRepo := gorm.NewPriceRepository()
	priceFounded, err := priceRepo.Find(priceId)
	assert.Nil(t, err)
	assert.NotNil(t, priceFounded)
	assert.Equal(t, priceId, priceFounded.Id())

	assert.Nil(t, priceRepo.Delete(price))
	priceFounded, err = priceRepo.Find(priceId)
	assert.NotNil(t, err)
	assert.Nil(t, priceFounded)
}

func TestPriceRepository_Find(t *testing.T) {
	price := factories.NewPriceFactory()
	priceId := price.Id()
	assert.NotNil(t, priceId)

	priceRepo := gorm.NewPriceRepository()
	priceFounded, err := priceRepo.Find(priceId)
	assert.Nil(t, err)
	assert.NotNil(t, priceFounded)
	assert.Equal(t, priceId, priceFounded.Id())
}

func TestPriceRepository_Update(t *testing.T) {
	price := factories.NewPriceFactory()
	newPrice := gofakeit.Float64()
	priceRepo := gorm.NewPriceRepository()
	price.Price = newPrice
	priceUpdated, err := priceRepo.Update(price)

	assert.Nil(t, err)
	assert.NotNil(t, priceUpdated)

	priceFounded, err := priceRepo.Find(priceUpdated.Id())
	assert.Nil(t, err)
	assert.NotNil(t, priceFounded)
	assert.Equal(t, newPrice, priceFounded.Price)
}

func TestNewPriceRepository(t *testing.T) {
	repo := gorm.NewPriceRepository()
	assert.NotNil(t, repo)
	assert.IsType(t, gorm.PriceRepository{}, *repo)
}
