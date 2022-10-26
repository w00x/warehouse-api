package mappers

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"warehouse/infrastructure/repository/mappers"
	"warehouse/infrastructure/repository/models"
	"warehouse/tests/factories"
)

func TestFromStockDomainToModel(t *testing.T) {
	domain := factories.NewStockDomainFactory()
	model := models.NewStock(domain.Id(), mappers.FromItemDomainToModel(domain.Item), mappers.FromRackDomainToModel(domain.Rack), domain.Quantity, domain.OperationDate, domain.ExpirationDate, domain.ItemId, domain.RackId, domain.Comment)
	toModel := mappers.FromStockDomainToModel(domain)

	assert.Equal(t, model.ID, toModel.ID)
	assert.Equal(t, model.Quantity, toModel.Quantity)
}

func TestFromStockModelToDomain(t *testing.T) {
	domain := factories.NewStockDomainFactory()
	model := models.NewStock(domain.Id(), mappers.FromItemDomainToModel(domain.Item), mappers.FromRackDomainToModel(domain.Rack), domain.Quantity, domain.OperationDate, domain.ExpirationDate, domain.ItemId, domain.RackId, domain.Comment)

	toDomain := mappers.FromStockModelToDomain(model)

	assert.Equal(t, model.ID, toDomain.Id())
	assert.Equal(t, model.Quantity, toDomain.Quantity)
}

func TestNewStockListDomainFromModel(t *testing.T) {
	domain := factories.NewStockDomainFactory()
	model := models.NewStock(domain.Id(), mappers.FromItemDomainToModel(domain.Item), mappers.FromRackDomainToModel(domain.Rack), domain.Quantity, domain.OperationDate, domain.ExpirationDate, domain.ItemId, domain.RackId, domain.Comment)
	var models []models.Stock
	models = append(models, *model)

	toDomains := mappers.NewStockListDomainFromModel(&models)

	assert.Equal(t, model.ID, (*toDomains)[0].Id())
	assert.Equal(t, model.Quantity, (*toDomains)[0].Quantity)
}
