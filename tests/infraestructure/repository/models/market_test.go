package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"warehouse/infrastructure/repository/models"
	"warehouse/tests/factories"
)

func TestNewMarket(t *testing.T) {
	domain := factories.NewMarketDomainFactory()
	model := models.NewMarket(domain.Id(), domain.Name)

	assert.Equal(t, model.Name, domain.Name)
}
