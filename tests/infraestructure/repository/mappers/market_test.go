package mappers

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"warehouse/infrastructure/repository/mappers"
	"warehouse/infrastructure/repository/models"
	"warehouse/tests/factories"
)

func TestFromMarketDomainToModel(t *testing.T) {
	domain := factories.NewMarketDomainFactory()
	model := models.NewMarket(domain.Id(), domain.Name)
	toModel := mappers.FromMarketDomainToModel(domain)

	assert.Equal(t, model.ID, toModel.ID)
	assert.Equal(t, model.Name, toModel.Name)
}

func TestFromMarketModelToDomain(t *testing.T) {
	domain := factories.NewMarketDomainFactory()
	model := models.NewMarket(domain.Id(), domain.Name)

	toDomain := mappers.FromMarketModelToDomain(model)

	assert.Equal(t, model.ID, toDomain.Id())
	assert.Equal(t, model.Name, toDomain.Name)
}

func TestNewMarketListDomainFromModel(t *testing.T) {
	domain := factories.NewMarketDomainFactory()
	model := models.NewMarket(domain.Id(), domain.Name)
	var models []models.Market
	models = append(models, *model)

	toDomains := mappers.NewMarketListDomainFromModel(&models)

	assert.Equal(t, model.ID, (*toDomains)[0].Id())
	assert.Equal(t, model.Name, (*toDomains)[0].Name)
}
