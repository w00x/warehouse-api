package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"warehouse/domain"
	"warehouse/tests/factories"
)

func TestItem_Id(t *testing.T) {
	domain := factories.NewItemDomainFactory()

	assert.IsType(t, "string", domain.Id())
	assert.Len(t, domain.Id(), 36)
}

func TestNewItem(t *testing.T) {
	object := factories.NewItemObjectFactory()
	domain := domain.NewItem("", object["name"].(string), object["unit_size_presentation"].(string),
		object["size_presentation"].(int), object["code"].(string), object["container"].(string), object["photo"].(string))

	assert.Equal(t, domain.Name, object["name"].(string))
}
