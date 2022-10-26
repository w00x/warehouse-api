package mappers

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"warehouse/infrastructure/repository/mappers"
	"warehouse/infrastructure/repository/models"
	"warehouse/tests/factories"
)

func TestFromItemDomainToModel(t *testing.T) {
	domain := factories.NewItemDomainFactory()
	model := models.NewItem(domain.Id(), domain.Name, domain.UnitSizePresentation, domain.SizePresentation, domain.Code, domain.Container, domain.Photo)
	toModel := mappers.FromItemDomainToModel(domain)

	assert.Equal(t, model.ID, toModel.ID)
	assert.Equal(t, model.Name, toModel.Name)
}

func TestFromItemModelToDomain(t *testing.T) {
	domain := factories.NewItemDomainFactory()
	model := models.NewItem(domain.Id(), domain.Name, domain.UnitSizePresentation, domain.SizePresentation, domain.Code, domain.Container, domain.Photo)

	toDomain := mappers.FromItemModelToDomain(model)

	assert.Equal(t, model.ID, toDomain.Id())
	assert.Equal(t, model.Name, toDomain.Name)
}

func TestNewItemListDomainFromModel(t *testing.T) {
	domain := factories.NewItemDomainFactory()
	model := models.NewItem(domain.Id(), domain.Name, domain.UnitSizePresentation, domain.SizePresentation, domain.Code, domain.Container, domain.Photo)
	var models []models.Item
	models = append(models, *model)

	toDomains := mappers.NewItemListDomainFromModel(&models)

	assert.Equal(t, model.ID, (*toDomains)[0].Id())
	assert.Equal(t, model.Name, (*toDomains)[0].Name)
}
