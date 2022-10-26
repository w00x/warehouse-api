package repository

import (
	"github.com/stretchr/testify/assert"
	"testing"
	repository2 "warehouse/infrastructure/factory/repository"
	"warehouse/infrastructure/repository/gorm"
)

func TestItemFactory(t *testing.T) {
	adapter := "gorm"
	repository, error := repository2.ItemFactory(adapter)

	assert.Nil(t, error)
	assert.IsType(t, gorm.NewItemRepository(), repository)

	adapter = "foo"
	repository, error = repository2.ItemFactory(adapter)

	assert.NotNil(t, error)
	assert.IsType(t, nil, repository)
}
