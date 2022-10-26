package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"warehouse/infrastructure/repository/models"
	"warehouse/tests/factories"
)

func TestNewInventory(t *testing.T) {
	domain := factories.NewInventoryDomainFactory()
	model := models.NewInventory(domain.Id(), domain.OperationDate)

	assert.Equal(t, model.OperationDate, domain.OperationDate)
}
