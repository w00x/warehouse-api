package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"warehouse/domain"
	"warehouse/tests/factories"
)

func TestRack_Id(t *testing.T) {
	domain := factories.NewRackDomainFactory()

	assert.IsType(t, "string", domain.Id())
	assert.Len(t, domain.Id(), 36)
}

func TestNewRack(t *testing.T) {
	object := factories.NewRackObjectFactory()
	domain := domain.NewRack("", object["name"].(string), object["code"].(string))

	assert.Equal(t, domain.Name, object["name"].(string))
}
