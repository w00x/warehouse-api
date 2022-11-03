package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"sort"
	"testing"
	"time"
	"warehouse/domain"
	"warehouse/infrastructure/repository/gorm"
	"warehouse/shared"
	"warehouse/tests/factories"
)

func TestStockIndexController(t *testing.T) {
	sizeOfStocks := 5
	stocks := factories.NewStockFactoryList(sizeOfStocks)

	resp, _ := http.Get(fmt.Sprintf("%s/v1/stock", Server().URL))

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	response := ParseResponseArray(resp)
	var ids []string

	for _, responseStock := range response {
		ids = append(ids, responseStock["id"].(string))
	}

	assert.Contains(t, ids, stocks[0].Id())
}

func TestStockGetController(t *testing.T) {
	stock := factories.NewStockFactory()

	resp, _ := http.Get(fmt.Sprintf("%s/v1/stock/%s", Server().URL, stock.Id()))

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	response := ParseResponse(resp)

	assert.Equal(t, stock.Id(), response["id"])
}

func TestStockCreateController(t *testing.T) {
	values := factories.NewStockObjectForCreateFactory()
	jsonData, err := json.Marshal(values)
	if err != nil {
		log.Fatal(err)
	}

	resp, _ := http.Post(fmt.Sprintf("%s/v1/stock", Server().URL), "application/json", bytes.NewBuffer(jsonData))

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	response := ParseResponse(resp)

	assert.Equal(t, values["quantity"], int(response["quantity"].(float64)))
}

func TestStockAllByInventoryController(t *testing.T) {
	sizeOfStocks := 5
	stocks := factories.NewStockFactoryList(sizeOfStocks)
	var listDates []shared.DateTime

	for _, stock := range stocks {
		listDates = append(listDates, stock.OperationDate)
	}

	sort.Sort(shared.DateTimeList(listDates))

	firstDate := listDates[0]
	inventoryDate := time.Time(firstDate).Add(-time.Hour * 24)
	inventoryDomain := domain.NewInventory("", shared.DateTime(inventoryDate))
	inventoryRepo := gorm.NewInventoryRepository()
	inventory, errCreate := inventoryRepo.Create(inventoryDomain)
	assert.Nil(t, errCreate)

	resp, _ := http.Get(fmt.Sprintf("%s/v1/stock/inventory/%s", Server().URL, inventory.Id()))

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	response := ParseResponseArray(resp)
	var ids []string

	for _, responseStock := range response {
		ids = append(ids, responseStock["id"].(string))
	}

	assert.Contains(t, ids, stocks[0].Id())
	assert.Contains(t, ids, stocks[1].Id())
	assert.Contains(t, ids, stocks[2].Id())
}
