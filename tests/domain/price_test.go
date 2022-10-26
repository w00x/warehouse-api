package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"warehouse/domain"
	"warehouse/shared"
	"warehouse/tests/factories"
)

func TestPrice_Id(t *testing.T) {
	domain := factories.NewPriceDomainFactory()

	assert.IsType(t, "string", domain.Id())
	assert.Len(t, domain.Id(), 36)
}

func TestNewPrice(t *testing.T) {
	object := factories.NewPriceObjectFactory()
	market := factories.NewMarketDomainFactory()
	item := factories.NewItemDomainFactory()
	date, _ := shared.StringToDate(object["date"].(string))
	domain := domain.NewPrice("", market,
		item, object["price"].(float64), *date, item.Id(),
		market.Id())

	assert.Equal(t, domain.Price, object["price"].(float64))
}
