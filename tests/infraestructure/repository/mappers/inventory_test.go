package mappers

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"warehouse/infrastructure/repository/mappers"
	"warehouse/infrastructure/repository/models"
	"warehouse/tests/factories"
)

func TestFromInventoryDomainToModel(t *testing.T) {
	domain := factories.NewInventoryDomainFactory()
	model := models.NewInventory(domain.Id(), domain.OperationDate)
	toModel := mappers.FromInventoryDomainToModel(domain)

	assert.Equal(t, model.ID, toModel.ID)
	assert.Equal(t, model.OperationDate, toModel.OperationDate)
}

func TestFromInventoryModelToDomain(t *testing.T) {
	domain := factories.NewInventoryDomainFactory()
	model := models.NewInventory(domain.Id(), domain.OperationDate)

	toDomain := mappers.FromInventoryModelToDomain(model)

	assert.Equal(t, model.ID, toDomain.Id())
	assert.Equal(t, model.OperationDate, toDomain.OperationDate)
}

func TestNewInventoryListDomainFromModel(t *testing.T) {
	domain := factories.NewInventoryDomainFactory()
	model := models.NewInventory(domain.Id(), domain.OperationDate)
	var models []models.Inventory
	models = append(models, *model)

	toDomains := mappers.NewInventoryListDomainFromModel(&models)

	assert.Equal(t, model.ID, (*toDomains)[0].Id())
	assert.Equal(t, model.OperationDate, (*toDomains)[0].OperationDate)
}
