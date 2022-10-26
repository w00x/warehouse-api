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

func TestPriceIndexController(t *testing.T) {
	sizeOfPrices := 5
	prices := factories.NewPriceFactoryList(sizeOfPrices)

	resp, _ := http.Get(fmt.Sprintf("%s/v1/price", Server().URL))

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	response := ParseResponseArray(resp)
	var ids []string

	for _, responsePrice := range response {
		ids = append(ids, responsePrice["id"].(string))
	}

	assert.Contains(t, ids, prices[0].Id())
}

func TestPriceGetController(t *testing.T) {
	price := factories.NewPriceFactory()

	resp, _ := http.Get(fmt.Sprintf("%s/v1/price/%s", Server().URL, price.Id()))

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	response := ParseResponse(resp)

	assert.Equal(t, price.Id(), response["id"])
}

func TestPriceCreateController(t *testing.T) {
	values := factories.NewPriceObjectForCreateFactory(t)
	jsonData, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}

	resp, _ := http.Post(fmt.Sprintf("%s/v1/price", Server().URL), "application/json", bytes.NewBuffer(jsonData))

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	response := ParseResponse(resp)

	assert.Equal(t, values["price"], response["price"])
}

func TestPriceUpdateController(t *testing.T) {
	price := factories.NewPriceFactory()

	values := factories.NewPriceObjectFactory()
	jsonData, _ := json.Marshal(values)

	req, _ := http.NewRequest(http.MethodPatch, fmt.Sprintf("%s/v1/price/%s", Server().URL, price.Id()), bytes.NewBuffer(jsonData))
	client := &http.Client{}
	resp, _ := client.Do(req)

	response := ParseResponse(resp)

	assert.Equal(t, values["price"], response["price"])
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestPriceDeleteController(t *testing.T) {
	price := factories.NewPriceFactory()

	req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/v1/price/%s", Server().URL, price.Id()), nil)
	client := &http.Client{}
	resp, _ := client.Do(req)

	_, err := gorm.NewPriceRepository().Find(price.Id())
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusNotFound, err.HttpStatusCode())
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}
