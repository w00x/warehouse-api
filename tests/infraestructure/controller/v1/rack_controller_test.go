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
	"testing"
	"warehouse/infrastructure/repository/gorm"
	"warehouse/tests/factories"
)

func TestRackIndexController(t *testing.T) {
	sizeOfRacks := 5
	racks := factories.NewRackFactoryList(sizeOfRacks)

	resp, _ := http.Get(fmt.Sprintf("%s/v1/rack", Server().URL))

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	response := ParseResponseArray(resp)
	var ids []string

	for _, responseRack := range response {
		ids = append(ids, responseRack["id"].(string))
	}

	assert.Contains(t, ids, racks[0].Id())
}

func TestRackGetController(t *testing.T) {
	rack := factories.NewRackFactory()

	resp, _ := http.Get(fmt.Sprintf("%s/v1/rack/%s", Server().URL, rack.Id()))

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response gin.H
	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, &response)

	assert.Equal(t, rack.Id(), response["id"])
}

func TestRackCreateController(t *testing.T) {
	values := factories.NewRackObjectFactory()
	jsonData, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}

	resp, _ := http.Post(fmt.Sprintf("%s/v1/rack", Server().URL), "application/json", bytes.NewBuffer(jsonData))

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var response gin.H
	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, &response)

	assert.Equal(t, values["name"], response["name"])
	assert.Equal(t, values["code"], response["code"])
}

func TestRackUpdateController(t *testing.T) {
	rack := factories.NewRackFactory()

	values := factories.NewRackObjectFactory()
	jsonData, _ := json.Marshal(values)

	req, _ := http.NewRequest(http.MethodPatch, fmt.Sprintf("%s/v1/rack/%s", Server().URL, rack.Id()), bytes.NewBuffer(jsonData))
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
	rack := factories.NewRackFactory()

	req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/v1/rack/%s", Server().URL, rack.Id()), nil)
	client := &http.Client{}
	resp, _ := client.Do(req)

	_, err := gorm.NewRackRepository().Find(rack.Id())
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusNotFound, err.HttpStatusCode())
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}
