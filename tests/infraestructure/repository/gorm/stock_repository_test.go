package gorm

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
	"time"
	"warehouse/domain"
	"warehouse/infrastructure/repository/gorm"
	"warehouse/shared"
	"warehouse/tests/factories"
)

func TestStockRepository_All(t *testing.T) {
	sizeOfStocks := 5
	stocks := factories.NewStockFactoryList(sizeOfStocks)
	stockRepo := gorm.NewStockRepository()
	allStocks, err := stockRepo.All()
	assert.Nil(t, err)

	var stocksIds []string
	for _, stock := range *allStocks {
		stocksIds = append(stocksIds, stock.Id())
	}

	assert.Contains(t, stocksIds, stocks[0].Id())
	assert.Contains(t, stocksIds, stocks[1].Id())
	assert.Contains(t, stocksIds, stocks[2].Id())
}

func TestStockRepository_Create(t *testing.T) {
	stockRepo := gorm.NewStockRepository()
	stockData := factories.NewStockDomainFactory()

	stock, err := stockRepo.Create(stockData)
	assert.Nil(t, err)

	assert.Equal(t, stock.Quantity, stockData.Quantity)
	assert.NotNil(t, stock.Id())
}

func TestStockRepository_Delete(t *testing.T) {
	stock := factories.NewStockFactory()
	stockId := stock.Id()
	assert.NotNil(t, stockId)

	stockRepo := gorm.NewStockRepository()
	stockFounded, err := stockRepo.Find(stockId)
	assert.Nil(t, err)
	assert.NotNil(t, stockFounded)
	assert.Equal(t, stockId, stockFounded.Id())

	assert.Nil(t, stockRepo.Delete(stock))
	stockFounded, err = stockRepo.Find(stockId)
	assert.NotNil(t, err)
	assert.Nil(t, stockFounded)
}

func TestStockRepository_Find(t *testing.T) {
	stock := factories.NewStockFactory()
	stockId := stock.Id()
	assert.NotNil(t, stockId)

	stockRepo := gorm.NewStockRepository()
	stockFounded, err := stockRepo.Find(stockId)
	assert.Nil(t, err)
	assert.NotNil(t, stockFounded)
	assert.Equal(t, stockId, stockFounded.Id())
}

func TestStockRepository_Update(t *testing.T) {
	stock := factories.NewStockFactory()
	newQuantity := gofakeit.Int32()
	stockRepo := gorm.NewStockRepository()
	stock.Quantity = int(newQuantity)
	stockUpdated, err := stockRepo.Update(stock)

	assert.Nil(t, err)
	assert.NotNil(t, stockUpdated)

	stockFounded, err := stockRepo.Find(stockUpdated.Id())
	assert.Nil(t, err)
	assert.NotNil(t, stockFounded)
	assert.Equal(t, int(newQuantity), stockFounded.Quantity)
}

func TestStockRepository_AllByInventory(t *testing.T) {
	sizeOfStocks := 5
	stocks := factories.NewStockFactoryList(sizeOfStocks)
	var listDates []shared.DateTime

	for _, stock := range stocks {
		listDates = append(listDates, stock.OperationDate)
	}

	sort.Sort(shared.DateTimeList(listDates))

	firstDate := listDates[0]
	inventoryDate := firstDate.Time.Add(-time.Hour * 24)
	inventoryDomain := domain.NewInventory("", shared.TimeToDateTime(inventoryDate))
	inventoryRepo := gorm.NewInventoryRepository()
	inventory, errCreate := inventoryRepo.Create(inventoryDomain)
	assert.Nil(t, errCreate)

	stockRepo := gorm.NewStockRepository()
	allStocks, err := stockRepo.AllByInventory(inventory.Id())
	assert.Nil(t, err)

	var stocksIds []string
	for _, stock := range *allStocks {
		stocksIds = append(stocksIds, stock.Id())
	}

	assert.Contains(t, stocksIds, stocks[0].Id())
	assert.Contains(t, stocksIds, stocks[1].Id())
	assert.Contains(t, stocksIds, stocks[2].Id())
}

func TestNewStockRepository(t *testing.T) {
	repo := gorm.NewStockRepository()
	assert.NotNil(t, repo)
	assert.IsType(t, gorm.StockRepository{}, *repo)
}
