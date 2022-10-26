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

func TestItemIndexController(t *testing.T) {
	sizeOfItems := 5
	items := factories.NewItemFactoryList(sizeOfItems)

	resp, _ := http.Get(fmt.Sprintf("%s/v1/item", Server().URL))

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	response := ParseResponseArray(resp)
	var ids []string

	for _, responseItem := range response {
		ids = append(ids, responseItem["id"].(string))
	}

	assert.Contains(t, ids, items[0].Id())
}

func TestItemGetController(t *testing.T) {
	item := factories.NewItemFactory()

	resp, _ := http.Get(fmt.Sprintf("%s/v1/item/%s", Server().URL, item.Id()))

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	response := ParseResponse(resp)

	assert.Equal(t, item.Id(), response["id"])
}

func TestItemCreateController(t *testing.T) {
	values := factories.NewItemObjectFactory()
	jsonData, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}

	resp, _ := http.Post(fmt.Sprintf("%s/v1/item", Server().URL), "application/json", bytes.NewBuffer(jsonData))

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	response := ParseResponse(resp)

	assert.Equal(t, values["name"], response["name"])
	assert.Equal(t, values["code"], response["code"])
}

func TestItemUpdateController(t *testing.T) {
	item := factories.NewItemFactory()

	values := factories.NewItemObjectFactory()
	jsonData, _ := json.Marshal(values)

	req, _ := http.NewRequest(http.MethodPatch, fmt.Sprintf("%s/v1/item/%s", Server().URL, item.Id()), bytes.NewBuffer(jsonData))
	client := &http.Client{}
	resp, _ := client.Do(req)

	response := ParseResponse(resp)

	assert.Equal(t, values["name"], response["name"])
	assert.Equal(t, values["code"], response["code"])
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestItemDeleteController(t *testing.T) {
	item := factories.NewItemFactory()

	req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/v1/item/%s", Server().URL, item.Id()), nil)
	client := &http.Client{}
	resp, _ := client.Do(req)

	_, err := gorm.NewItemRepository().Find(item.Id())
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusNotFound, err.HttpStatusCode())
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}
