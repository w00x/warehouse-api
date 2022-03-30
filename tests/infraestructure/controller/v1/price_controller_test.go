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

func TestPriceIndexController(t *testing.T) {
	sizeOfPrices := 5
	prices := factories.NewPriceFactoryList(sizeOfPrices, t)
	router := infraestructure.Routes()
	server := httptest.NewServer(router)

	resp, _ := http.Get(fmt.Sprintf("%s/v1/price", server.URL))

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response []gin.H
	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, &response)
	var ids []float64

	for _, responsePrice := range response {
		ids = append(ids, responsePrice["id"].(float64))
	}

	assert.Contains(t, ids, float64(prices[0].Id))
}

func TestPriceGetController(t *testing.T) {
	price := factories.NewPriceFactory(t)
	router := infraestructure.Routes()
	server := httptest.NewServer(router)

	resp, _ := http.Get(fmt.Sprintf("%s/v1/price/%d", server.URL, price.Id))

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response gin.H
	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, &response)

	assert.Equal(t, float64(price.Id), response["id"])
}

func TestPriceCreateController(t *testing.T) {
	defer factories.CleanPrice()
	router := infraestructure.Routes()
	server := httptest.NewServer(router)

	values := factories.NewPriceObjectForCreateFactory(t)
	jsonData, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}

	resp, _ := http.Post(fmt.Sprintf("%s/v1/price", server.URL), "application/json", bytes.NewBuffer(jsonData))

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var response gin.H
	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, &response)

	assert.Equal(t, values["price"], response["price"])
}

func TestPriceUpdateController(t *testing.T) {
	price := factories.NewPriceFactory(t)
	router := infraestructure.Routes()
	server := httptest.NewServer(router)

	values := factories.NewPriceObjectFactory()
	jsonData, _ := json.Marshal(values)

	req, _ := http.NewRequest(http.MethodPatch, fmt.Sprintf("%s/v1/price/%d", server.URL, price.Id), bytes.NewBuffer(jsonData))
	client := &http.Client{}
	resp, _ := client.Do(req)

	var response gin.H
	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, &response)

	assert.Equal(t, values["price"], response["price"])
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestPriceDeleteController(t *testing.T) {
	price := factories.NewPriceFactory(t)
	router := infraestructure.Routes()
	server := httptest.NewServer(router)

	req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/v1/price/%d", server.URL, price.Id), nil)
	client := &http.Client{}
	resp, _ := client.Do(req)

	_, err := postgres.NewPriceRepository().Find(price.Id)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusNotFound, err.HttpStatusCode())
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}