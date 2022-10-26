package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"warehouse/domain"
	"warehouse/shared"
	"warehouse/tests/factories"
)

func TestInventory_Id(t *testing.T) {
	domain := factories.NewInventoryDomainFactory()

	assert.IsType(t, "string", domain.Id())
	assert.Len(t, domain.Id(), 36)
}

func TestNewInventory(t *testing.T) {
	object := factories.NewInventoryObjectFactory()
	date, _ := shared.StringToDate(object["operation_date"].(string))
	domain := domain.NewInventory("", *date)

	assert.Equal(t, domain.OperationDate, *date)
}
