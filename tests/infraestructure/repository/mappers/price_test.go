package mappers

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"warehouse/infrastructure/repository/mappers"
	"warehouse/infrastructure/repository/models"
	"warehouse/tests/factories"
)

func TestFromPriceDomainToModel(t *testing.T) {
	domain := factories.NewPriceDomainFactory()
	model := models.NewPrice(domain.Id(), mappers.FromMarketDomainToModel(domain.Market), mappers.FromItemDomainToModel(domain.Item), domain.Price, domain.Date, domain.ItemId, domain.MarketId)
	toModel := mappers.FromPriceDomainToModel(domain)

	assert.Equal(t, model.ID, toModel.ID)
	assert.Equal(t, model.Price, toModel.Price)
}

func TestFromPriceModelToDomain(t *testing.T) {
	domain := factories.NewPriceDomainFactory()
	model := models.NewPrice(domain.Id(), mappers.FromMarketDomainToModel(domain.Market), mappers.FromItemDomainToModel(domain.Item), domain.Price, domain.Date, domain.ItemId, domain.MarketId)

	toDomain := mappers.FromPriceModelToDomain(model)

	assert.Equal(t, model.ID, toDomain.Id())
	assert.Equal(t, model.Price, toDomain.Price)
}

func TestNewPriceListDomainFromModel(t *testing.T) {
	domain := factories.NewPriceDomainFactory()
	model := models.NewPrice(domain.Id(), mappers.FromMarketDomainToModel(domain.Market), mappers.FromItemDomainToModel(domain.Item), domain.Price, domain.Date, domain.ItemId, domain.MarketId)
	var models []models.Price
	models = append(models, *model)

	toDomains := mappers.NewPriceListDomainFromModel(&models)

	assert.Equal(t, model.ID, (*toDomains)[0].Id())
	assert.Equal(t, model.Price, (*toDomains)[0].Price)
}
