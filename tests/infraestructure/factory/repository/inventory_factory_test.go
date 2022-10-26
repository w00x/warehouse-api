package repository

import (
	"github.com/stretchr/testify/assert"
	"testing"
	repository2 "warehouse/infrastructure/factory/repository"
	"warehouse/infrastructure/repository/gorm"
)

func TestInventoryFactory(t *testing.T) {
	adapter := "gorm"
	repository, error := repository2.InventoryFactory(adapter)

	assert.Nil(t, error)
	assert.IsType(t, gorm.NewInventoryRepository(), repository)

	adapter = "foo"
	repository, error = repository2.InventoryFactory(adapter)

	assert.NotNil(t, error)
	assert.IsType(t, nil, repository)
}
