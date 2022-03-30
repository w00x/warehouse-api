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

func TestItemIndexController(t *testing.T) {
	sizeOfItems := 5
	items := factories.NewItemFactoryList(sizeOfItems, t)
	router := infraestructure.Routes()
	server := httptest.NewServer(router)

	resp, _ := http.Get(fmt.Sprintf("%s/v1/item", server.URL))

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response []gin.H
	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, &response)
	var ids []float64

	for _, responseItem := range response {
		ids = append(ids, responseItem["id"].(float64))
	}

	assert.Contains(t, ids, float64(items[0].Id))
}

func TestItemGetController(t *testing.T) {
	item := factories.NewItemFactory(t)
	router := infraestructure.Routes()
	server := httptest.NewServer(router)

	resp, _ := http.Get(fmt.Sprintf("%s/v1/item/%d", server.URL, item.Id))

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response gin.H
	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, &response)

	assert.Equal(t, float64(item.Id), response["id"])
}

func TestItemCreateController(t *testing.T) {
	defer factories.CleanItem()
	router := infraestructure.Routes()
	server := httptest.NewServer(router)

	values := factories.NewItemObjectFactory()
	jsonData, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}

	resp, _ := http.Post(fmt.Sprintf("%s/v1/item", server.URL), "application/json", bytes.NewBuffer(jsonData))

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var response gin.H
	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, &response)

	assert.Equal(t, values["name"], response["name"])
	assert.Equal(t, values["code"], response["code"])
}

func TestItemUpdateController(t *testing.T) {
	item := factories.NewItemFactory(t)
	router := infraestructure.Routes()
	server := httptest.NewServer(router)

	values := factories.NewItemObjectFactory()
	jsonData, _ := json.Marshal(values)

	req, _ := http.NewRequest(http.MethodPatch, fmt.Sprintf("%s/v1/item/%d", server.URL, item.Id), bytes.NewBuffer(jsonData))
	client := &http.Client{}
	resp, _ := client.Do(req)

	var response gin.H
	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, &response)

	assert.Equal(t, values["name"], response["name"])
	assert.Equal(t, values["code"], response["code"])
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestItemDeleteController(t *testing.T) {
	item := factories.NewItemFactory(t)
	router := infraestructure.Routes()
	server := httptest.NewServer(router)

	req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/v1/item/%d", server.URL, item.Id), nil)
	client := &http.Client{}
	resp, _ := client.Do(req)

	_, err := postgres.NewItemRepository().Find(item.Id)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusNotFound, err.HttpStatusCode())
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}