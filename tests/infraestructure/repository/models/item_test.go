package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"warehouse/infrastructure/repository/models"
	"warehouse/tests/factories"
)

func TestNewItem(t *testing.T) {
	domain := factories.NewItemDomainFactory()
	model := models.NewItem(domain.Id(), domain.Name, domain.UnitSizePresentation,
		domain.SizePresentation, domain.Code, domain.Container, domain.Photo)

	assert.Equal(t, model.Name, domain.Name)
}
