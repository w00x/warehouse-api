package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"testing"
	"warehouse/infrastructure/repository/gorm"
	"warehouse/tests/factories"
)

func TestInventoryIndexController(t *testing.T) {
	sizeOfInventories := 5
	inventories := factories.NewInventoryFactoryList(sizeOfInventories)
	resp, _ := http.Get(fmt.Sprintf("%s/v1/inventory", Server().URL))

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	response := ParseResponseArray(resp)

	var ids []string
	for _, responseInventory := range response {
		ids = append(ids, responseInventory["id"].(string))
	}

	assert.Contains(t, ids, inventories[0].Id())
}

func TestInventoryGetController(t *testing.T) {
	inventory := factories.NewInventoryFactory()
	resp, _ := http.Get(fmt.Sprintf("%s/v1/inventory/%s", Server().URL, inventory.Id()))

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	response := ParseResponse(resp)

	assert.Equal(t, inventory.Id(), response["id"].(string))
}

func TestInventoryCreateController(t *testing.T) {
	date := "2022-03-05 12:58:00"
	values := map[string]interface{}{"operation_date": date}
	jsonData, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}

	resp, _ := http.Post(fmt.Sprintf("%s/v1/inventory", Server().URL), "application/json", bytes.NewBuffer(jsonData))

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	response := ParseResponse(resp)

	assert.Equal(t, date, response["operation_date"])
}

func TestInventoryUpdateController(t *testing.T) {
	inventory := factories.NewInventoryFactory()
	values := factories.NewInventoryObjectFactory()
	jsonData, _ := json.Marshal(values)

	req, _ := http.NewRequest(http.MethodPatch, fmt.Sprintf("%s/v1/inventory/%s", Server().URL, inventory.Id()), bytes.NewBuffer(jsonData))
	client := &http.Client{}
	resp, _ := client.Do(req)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	response := ParseResponse(resp)
	assert.Equal(t, values["operation_date"], response["operation_date"])
}

func TestInventoryDeleteController(t *testing.T) {
	inventory := factories.NewInventoryFactory()
	req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/v1/inventory/%s", Server().URL, inventory.Id()), nil)
	client := &http.Client{}
	resp, _ := client.Do(req)

	_, err := gorm.NewInventoryRepository().Find(inventory.Id())
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusNotFound, err.HttpStatusCode())
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}
