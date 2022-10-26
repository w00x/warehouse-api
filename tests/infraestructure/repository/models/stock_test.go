package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"warehouse/infrastructure/repository/mappers"
	"warehouse/infrastructure/repository/models"
	"warehouse/tests/factories"
)

func TestNewStock(t *testing.T) {
	domain := factories.NewStockDomainFactory()
	model := models.NewStock(domain.Id(), mappers.FromItemDomainToModel(domain.Item),
		mappers.FromRackDomainToModel(domain.Rack), domain.Quantity, domain.OperationDate,
		domain.ExpirationDate, domain.Comment, domain.ItemId, domain.RackId)

	assert.Equal(t, model.Quantity, domain.Quantity)
}
