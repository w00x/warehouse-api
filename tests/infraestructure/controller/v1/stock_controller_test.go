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

func TestStockIndexController(t *testing.T) {
	sizeOfStocks := 5
	stocks := factories.NewStockFactoryList(sizeOfStocks, t)
	router := infraestructure.Routes()
	server := httptest.NewServer(router)

	resp, _ := http.Get(fmt.Sprintf("%s/v1/stock", server.URL))

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response []gin.H
	data, _ := ioutil.ReadAll(resp.Body)
	err := json.Unmarshal(data, &response)
	if err != nil {
		log.Fatal(err)
	}
	var ids []float64

	for _, responseStock := range response {
		ids = append(ids, responseStock["id"].(float64))
	}

	assert.Contains(t, ids, float64(stocks[0].Id))
}

func TestStockGetController(t *testing.T) {
	stock := factories.NewStockFactory(t)
	router := infraestructure.Routes()
	server := httptest.NewServer(router)

	resp, _ := http.Get(fmt.Sprintf("%s/v1/stock/%d", server.URL, stock.Id))

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response gin.H
	data, _ := ioutil.ReadAll(resp.Body)
	err := json.Unmarshal(data, &response)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, float64(stock.Id), response["id"])
}

func TestStockCreateController(t *testing.T) {
	defer factories.CleanStock()
	router := infraestructure.Routes()
	server := httptest.NewServer(router)

	values := factories.NewStockObjectForCreateFactory(t)
	jsonData, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}

	resp, _ := http.Post(fmt.Sprintf("%s/v1/stock", server.URL), "application/json", bytes.NewBuffer(jsonData))

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var response gin.H
	data, _ := ioutil.ReadAll(resp.Body)
	errUn := json.Unmarshal(data, &response)
	if err != nil {
		log.Fatal(errUn)
	}

	assert.Equal(t, values["quantity"], int(response["quantity"].(float64)))
}

func TestStockUpdateController(t *testing.T) {
	stock := factories.NewStockFactory(t)
	router := infraestructure.Routes()
	server := httptest.NewServer(router)

	values := factories.NewStockObjectFactory()
	jsonData, _ := json.Marshal(values)

	req, _ := http.NewRequest(http.MethodPatch, fmt.Sprintf("%s/v1/stock/%d", server.URL, stock.Id), bytes.NewBuffer(jsonData))
	client := &http.Client{}
	resp, _ := client.Do(req)

	var response gin.H
	data, _ := ioutil.ReadAll(resp.Body)
	err := json.Unmarshal(data, &response)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, values["quantity"], int(response["quantity"].(float64)))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestStockDeleteController(t *testing.T) {
	stock := factories.NewStockFactory(t)
	router := infraestructure.Routes()
	server := httptest.NewServer(router)

	req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/v1/stock/%d", server.URL, stock.Id), nil)
	client := &http.Client{}
	resp, _ := client.Do(req)

	_, err := postgres.NewStockRepository().Find(stock.Id)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusNotFound, err.HttpStatusCode())
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}