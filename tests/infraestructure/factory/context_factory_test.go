package factory

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
	adapters "warehouse/infrastructure/controller/context"
	"warehouse/infrastructure/factory"
)

func TestContextFactory(t *testing.T) {
	adapter := "gin"
	contextParam := gin.Context{
		Request:  nil,
		Writer:   nil,
		Params:   nil,
		Keys:     nil,
		Errors:   nil,
		Accepted: nil,
	}
	context, error := factory.ContextFactory(adapter, &contextParam)

	assert.Nil(t, error)
	assert.IsType(t, &adapters.GinContextAdapter{Ctx: &contextParam}, context)

	adapter = "foo"
	context, error = factory.ContextFactory(adapter, &contextParam)

	assert.NotNil(t, error)
	assert.IsType(t, nil, context)
}
