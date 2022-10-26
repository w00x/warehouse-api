package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"warehouse/infrastructure/repository/mappers"
	"warehouse/infrastructure/repository/models"
	"warehouse/tests/factories"
)

func TestNewPrice(t *testing.T) {
	domain := factories.NewPriceDomainFactory()
	model := models.NewPrice(domain.Id(), mappers.FromMarketDomainToModel(domain.Market),
		mappers.FromItemDomainToModel(domain.Item), domain.Price, domain.Date, domain.ItemId,
		domain.MarketId)

	assert.Equal(t, model.Price, domain.Price)
}
