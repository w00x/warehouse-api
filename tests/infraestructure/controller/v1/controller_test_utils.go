package v1

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"warehouse/infrastructure"
)

const FactoryAdapter = "gorm"

func Server() *httptest.Server {
	router := infrastructure.Routes(FactoryAdapter)
	return httptest.NewServer(router)
}

func ParseResponseArray(resp *http.Response) []gin.H {
	var response []gin.H
	data, errRead := ioutil.ReadAll(resp.Body)
	if errRead != nil {
		log.Fatal(errRead)
	}

	err := json.Unmarshal(data, &response)
	if err != nil {
		log.Fatal(err)
	}
	return response
}

func ParseResponse(resp *http.Response) gin.H {
	var response gin.H
	d := json.NewDecoder(resp.Body)
	if err := d.Decode(&response); err != nil {
		panic(err)
	}

	return response
}
