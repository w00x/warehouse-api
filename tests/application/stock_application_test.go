package application

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
	"time"
	"warehouse/application"
	"warehouse/domain"
	"warehouse/infrastructure/repository/gorm"
	"warehouse/shared"
	"warehouse/tests"
	"warehouse/tests/factories"
)

func TestStockApplication_All(t *testing.T) {
	tests.CleanStock()
	tests.CleanInventory()
	sizeOfInventories := 5
	stocksList := factories.NewStockFactoryList(sizeOfInventories)

	var dates []shared.DateTime
	for _, stock := range stocksList {
		dates = append(dates, stock.OperationDate)
	}

	datesList := shared.DateTimeList(dates)
	sort.Sort(&datesList)

	inventoryDate := datesList[1]

	repoInventory := gorm.NewInventoryRepository()
	repoInventory.Create(domain.NewInventory("", inventoryDate))

	repoStock := gorm.NewStockRepository()
	stockApplication := application.NewStockApplication(repoStock)
	stocks, err := stockApplication.All()

	assert.Nil(t, err)
	assert.NotEqual(t, len(stocksList), len(*stocks))
	assert.Equal(t, len(stocksList)-1, len(*stocks))
	tests.CleanStock()
	tests.CleanInventory()
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

func TestStockApplication_AllByInventory(t *testing.T) {
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

	repo := gorm.NewStockRepository()
	stockApplication := application.NewStockApplication(repo)
	stocksList, err := stockApplication.AllByInventory(inventory.Id())

	assert.Nil(t, err)
	var ids []string
	for _, stock := range *stocksList {
		ids = append(ids, stock.Id())
	}

	assert.Contains(t, ids, stocks[0].Id())
}

func TestNewStockApplication(t *testing.T) {
	repo := gorm.NewStockRepository()
	stockApplication := application.NewStockApplication(repo)

	assert.NotNil(t, stockApplication)
}
