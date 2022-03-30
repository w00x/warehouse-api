package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"warehouse/infraestructure"
	"warehouse/infraestructure/repository/postgres"
	"warehouse/tests/factories"
)

func TestInventoryIndexController(t *testing.T) {
	sizeOfInvenotories := 5
	inventories := factories.NewInventoryFactoryList(sizeOfInvenotories, t)
	router := infraestructure.Routes()
	server := httptest.NewServer(router)

	resp, _ := http.Get(fmt.Sprintf("%s/v1/inventory", server.URL))

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response []gin.H
	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, &response)
	var ids []float64

	for _, responseInventory := range response {
		ids = append(ids, responseInventory["id"].(float64))
	}

	assert.Contains(t, ids, float64(inventories[0].Id))
}

func TestInventoryGetController(t *testing.T) {
	inventory := factories.NewInventoryFactory(t)
	router := infraestructure.Routes()
	server := httptest.NewServer(router)

	resp, _ := http.Get(fmt.Sprintf("%s/v1/inventory/%d", server.URL, inventory.Id))

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response gin.H
	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, &response)

	assert.Equal(t, float64(inventory.Id), response["id"])
}

func TestInventoryCreateController(t *testing.T) {
	defer factories.CleanInventory()
	router := infraestructure.Routes()
	server := httptest.NewServer(router)

	date  := "2022-03-05 12:58:00"
	values := map[string]interface{}{"operation_date": date}
	jsonData, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}

	resp, _ := http.Post(fmt.Sprintf("%s/v1/inventory", server.URL), "application/json", bytes.NewBuffer(jsonData))

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var response gin.H
	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, &response)

	assert.Equal(t, date, response["operation_date"])
}

func TestInventoryUpdateController(t *testing.T) {
	inventory := factories.NewInventoryFactory(t)
	router := infraestructure.Routes()
	server := httptest.NewServer(router)

	values := factories.NewInventoryObjectFactory()
	jsonData, _ := json.Marshal(values)

	req, _ := http.NewRequest(http.MethodPatch, fmt.Sprintf("%s/v1/inventory/%d", server.URL, inventory.Id), bytes.NewBuffer(jsonData))
	client := &http.Client{}
	resp, _ := client.Do(req)

	var response gin.H
	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, &response)

	assert.Equal(t, values["operation_date"], response["operation_date"])
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestInventoryDeleteController(t *testing.T) {
	inventory := factories.NewInventoryFactory(t)
	router := infraestructure.Routes()
	server := httptest.NewServer(router)

	req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/v1/inventory/%d", server.URL, inventory.Id), nil)
	client := &http.Client{}
	resp, _ := client.Do(req)

	_, err := postgres.NewInventoryRepository().Find(inventory.Id)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusNotFound, err.HttpStatusCode())
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}