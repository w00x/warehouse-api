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

func TestMarketIndexController(t *testing.T) {
	sizeOfMarkets := 5
	markets := factories.NewMarketFactoryList(sizeOfMarkets, t)
	router := infraestructure.Routes()
	server := httptest.NewServer(router)

	resp, _ := http.Get(fmt.Sprintf("%s/v1/market", server.URL))

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response []gin.H
	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, &response)
	var ids []float64

	for _, responseMarket := range response {
		ids = append(ids, responseMarket["id"].(float64))
	}

	assert.Contains(t, ids, float64(markets[0].Id))
}

func TestMarketGetController(t *testing.T) {
	market := factories.NewMarketFactory(t)
	router := infraestructure.Routes()
	server := httptest.NewServer(router)

	resp, _ := http.Get(fmt.Sprintf("%s/v1/market/%d", server.URL, market.Id))

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response gin.H
	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, &response)

	assert.Equal(t, float64(market.Id), response["id"])
}

func TestMarketCreateController(t *testing.T) {
	defer factories.CleanMarket()
	router := infraestructure.Routes()
	server := httptest.NewServer(router)

	values := factories.NewMarketObjectFactory()
	jsonData, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}

	resp, _ := http.Post(fmt.Sprintf("%s/v1/market", server.URL), "application/json", bytes.NewBuffer(jsonData))

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var response gin.H
	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, &response)

	assert.Equal(t, values["name"], response["name"])
}

func TestMarketUpdateController(t *testing.T) {
	market := factories.NewMarketFactory(t)
	router := infraestructure.Routes()
	server := httptest.NewServer(router)

	values := factories.NewMarketObjectFactory()
	jsonData, _ := json.Marshal(values)

	req, _ := http.NewRequest(http.MethodPatch, fmt.Sprintf("%s/v1/market/%d", server.URL, market.Id), bytes.NewBuffer(jsonData))
	client := &http.Client{}
	resp, _ := client.Do(req)

	var response gin.H
	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, &response)

	assert.Equal(t, values["name"], response["name"])
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestMarketDeleteController(t *testing.T) {
	market := factories.NewMarketFactory(t)
	router := infraestructure.Routes()
	server := httptest.NewServer(router)

	req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/v1/market/%d", server.URL, market.Id), nil)
	client := &http.Client{}
	resp, _ := client.Do(req)

	_, err := postgres.NewMarketRepository().Find(market.Id)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusNotFound, err.HttpStatusCode())
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}