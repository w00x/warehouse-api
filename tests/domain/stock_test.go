package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"warehouse/domain"
	"warehouse/shared"
	"warehouse/tests/factories"
)

func TestStock_Id(t *testing.T) {
	domain := factories.NewStockDomainFactory()

	assert.IsType(t, "string", domain.Id())
	assert.Len(t, domain.Id(), 36)
}

func TestNewStock(t *testing.T) {
	object := factories.NewStockObjectFactory()
	rack := factories.NewRackDomainFactory()
	item := factories.NewItemDomainFactory()
	operationDate, _ := shared.StringToDate(object["operation_date"].(string))
	expirationDate, _ := shared.StringToDate(object["expiration_date"].(string))
	comment := object["comment"].(string)
	domain := domain.NewStock("", item,
		rack, object["quantity"].(int), *operationDate, *expirationDate, comment, item.Id(),
		rack.Id())

	assert.Equal(t, domain.Quantity, object["quantity"].(int))
}
