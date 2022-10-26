package repository

import (
	"github.com/stretchr/testify/assert"
	"testing"
	repository2 "warehouse/infrastructure/factory/repository"
	"warehouse/infrastructure/repository/gorm"
)

func TestMarketFactory(t *testing.T) {
	adapter := "gorm"
	repository, error := repository2.MarketFactory(adapter)

	assert.Nil(t, error)
	assert.IsType(t, gorm.NewMarketRepository(), repository)

	adapter = "foo"
	repository, error = repository2.MarketFactory(adapter)

	assert.NotNil(t, error)
	assert.IsType(t, nil, repository)
}
