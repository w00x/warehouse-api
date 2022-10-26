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

func TestMarketIndexController(t *testing.T) {
	sizeOfMarkets := 5
	markets := factories.NewMarketFactoryList(sizeOfMarkets)

	resp, _ := http.Get(fmt.Sprintf("%s/v1/market", Server().URL))

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	response := ParseResponseArray(resp)
	var ids []string

	for _, responseMarket := range response {
		ids = append(ids, responseMarket["id"].(string))
	}

	assert.Contains(t, ids, markets[0].Id())
}

func TestMarketGetController(t *testing.T) {
	market := factories.NewMarketFactory()

	resp, _ := http.Get(fmt.Sprintf("%s/v1/market/%s", Server().URL, market.Id()))

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	response := ParseResponse(resp)

	assert.Equal(t, market.Id(), response["id"])
}

func TestMarketCreateController(t *testing.T) {
	values := factories.NewMarketObjectFactory()
	jsonData, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}

	resp, _ := http.Post(fmt.Sprintf("%s/v1/market", Server().URL), "application/json", bytes.NewBuffer(jsonData))

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	response := ParseResponse(resp)

	assert.Equal(t, values["name"], response["name"])
}

func TestMarketUpdateController(t *testing.T) {
	market := factories.NewMarketFactory()

	values := factories.NewMarketObjectFactory()
	jsonData, _ := json.Marshal(values)

	req, _ := http.NewRequest(http.MethodPatch, fmt.Sprintf("%s/v1/market/%s", Server().URL, market.Id()), bytes.NewBuffer(jsonData))
	client := &http.Client{}
	resp, _ := client.Do(req)

	response := ParseResponse(resp)

	assert.Equal(t, values["name"], response["name"])
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestMarketDeleteController(t *testing.T) {
	market := factories.NewMarketFactory()

	req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/v1/market/%s", Server().URL, market.Id()), nil)
	client := &http.Client{}
	resp, _ := client.Do(req)

	_, err := gorm.NewMarketRepository().Find(market.Id())
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusNotFound, err.HttpStatusCode())
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}
