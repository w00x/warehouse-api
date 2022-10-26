package mappers

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"warehouse/infrastructure/repository/mappers"
	"warehouse/infrastructure/repository/models"
	"warehouse/tests/factories"
)

func TestFromRackDomainToModel(t *testing.T) {
	domain := factories.NewRackDomainFactory()
	model := models.NewRack(domain.Id(), domain.Name, domain.Code)
	toModel := mappers.FromRackDomainToModel(domain)

	assert.Equal(t, model.ID, toModel.ID)
	assert.Equal(t, model.Name, toModel.Name)
}

func TestFromRackModelToDomain(t *testing.T) {
	domain := factories.NewRackDomainFactory()
	model := models.NewRack(domain.Id(), domain.Name, domain.Code)

	toDomain := mappers.FromRackModelToDomain(model)

	assert.Equal(t, model.ID, toDomain.Id())
	assert.Equal(t, model.Name, toDomain.Name)
}

func TestNewRackListDomainFromModel(t *testing.T) {
	domain := factories.NewRackDomainFactory()
	model := models.NewRack(domain.Id(), domain.Name, domain.Code)
	var models []models.Rack
	models = append(models, *model)

	toDomains := mappers.NewRackListDomainFromModel(&models)

	assert.Equal(t, model.ID, (*toDomains)[0].Id())
	assert.Equal(t, model.Name, (*toDomains)[0].Name)
}
