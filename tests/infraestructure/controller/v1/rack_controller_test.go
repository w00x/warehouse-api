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

func TestRackIndexController(t *testing.T) {
	sizeOfRacks := 5
	racks := factories.NewRackFactoryList(sizeOfRacks, t)
	router := infraestructure.Routes()
	server := httptest.NewServer(router)

	resp, _ := http.Get(fmt.Sprintf("%s/v1/rack", server.URL))

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response []gin.H
	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, &response)
	var ids []float64

	for _, responseRack := range response {
		ids = append(ids, responseRack["id"].(float64))
	}

	assert.Contains(t, ids, float64(racks[0].Id))
}

func TestRackGetController(t *testing.T) {
	rack := factories.NewRackFactory(t)
	router := infraestructure.Routes()
	server := httptest.NewServer(router)

	resp, _ := http.Get(fmt.Sprintf("%s/v1/rack/%d", server.URL, rack.Id))

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response gin.H
	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, &response)

	assert.Equal(t, float64(rack.Id), response["id"])
}

func TestRackCreateController(t *testing.T) {
	defer factories.CleanRack()
	router := infraestructure.Routes()
	server := httptest.NewServer(router)

	values := factories.NewRackObjectFactory()
	jsonData, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}

	resp, _ := http.Post(fmt.Sprintf("%s/v1/rack", server.URL), "application/json", bytes.NewBuffer(jsonData))

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var response gin.H
	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, &response)

	assert.Equal(t, values["name"], response["name"])
	assert.Equal(t, values["code"], response["code"])
}

func TestRackUpdateController(t *testing.T) {
	rack := factories.NewRackFactory(t)
	router := infraestructure.Routes()
	server := httptest.NewServer(router)

	values := factories.NewRackObjectFactory()
	jsonData, _ := json.Marshal(values)

	req, _ := http.NewRequest(http.MethodPatch, fmt.Sprintf("%s/v1/rack/%d", server.URL, rack.Id), bytes.NewBuffer(jsonData))
	client := &http.Client{}
	resp, _ := client.Do(req)

	var response gin.H
	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, &response)

	assert.Equal(t, values["name"], response["name"])
	assert.Equal(t, values["code"], response["code"])
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestRackDeleteController(t *testing.T) {
	rack := factories.NewRackFactory(t)
	router := infraestructure.Routes()
	server := httptest.NewServer(router)

	req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/v1/rack/%d", server.URL, rack.Id), nil)
	client := &http.Client{}
	resp, _ := client.Do(req)

	_, err := postgres.NewRackRepository().Find(rack.Id)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusNotFound, err.HttpStatusCode())
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}