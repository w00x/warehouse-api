package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"warehouse/infrastructure/repository/models"
	"warehouse/tests/factories"
)

func TestNewRack(t *testing.T) {
	domain := factories.NewRackDomainFactory()
	model := models.NewRack(domain.Id(), domain.Name, domain.Code)

	assert.Equal(t, model.Name, domain.Name)
}
