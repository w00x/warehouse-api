package application

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"warehouse/application"
	"warehouse/infrastructure/repository/gorm"
	"warehouse/tests/factories"
)

func TestStockApplication_All(t *testing.T) {
	sizeOfInventories := 5
	stocksList := factories.NewStockFactoryList(sizeOfInventories)
	repo := gorm.NewStockRepository()
	stockApplication := application.NewStockApplication(repo)
	stocks, err := stockApplication.All()

	assert.Nil(t, err)
	var ids []string
	for _, stock := range stocksList {
		ids = append(ids, stock.Id())
	}

	assert.Contains(t, ids, (*stocks)[0].Id())
}

func TestStockApplication_Create(t *testing.T) {
	stock := factories.NewStockDomainFactory()
	repo := gorm.NewStockRepository()
	stockApplication := application.NewStockApplication(repo)
	newStock, err := stockApplication.Create(stock)

	assert.Nil(t, err)
	assert.NotNil(t, newStock.Id())
	assert.Equal(t, stock.Comment, newStock.Comment)
}

func TestStockApplication_Show(t *testing.T) {
	stock := factories.NewStockFactory()
	repo := gorm.NewStockRepository()
	stockApplication := application.NewStockApplication(repo)
	findedStock, err := stockApplication.Show(stock.Id())

	assert.Nil(t, err)
	assert.Equal(t, stock.Id(), findedStock.Id())
}

func TestNewStockApplication(t *testing.T) {
	repo := gorm.NewStockRepository()
	stockApplication := application.NewStockApplication(repo)

	assert.NotNil(t, stockApplication)
}
