package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"warehouse/domain"
	"warehouse/tests/factories"
)

func TestMarket_Id(t *testing.T) {
	domain := factories.NewMarketDomainFactory()

	assert.IsType(t, "string", domain.Id())
	assert.Len(t, domain.Id(), 36)
}

func TestNewMarket(t *testing.T) {
	object := factories.NewMarketObjectFactory()
	domain := domain.NewMarket("", object["name"].(string))

	assert.Equal(t, domain.Name, object["name"].(string))
}
