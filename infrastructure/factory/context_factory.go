package factory

import (
	"github.com/gin-gonic/gin"
	adapters "warehouse/infrastructure/controller/context"
	"warehouse/infrastructure/errors"
)

func ContextFactory(adapter string, ctx interface{}) (adapters.IContextAdapter, errors.IBaseError) {
	if adapter == "gin" {
		return &adapters.GinContextAdapter{Ctx: ctx.(*gin.Context)}, nil
	}

	return nil, errors.NewNotFoundError("Context Factory not found")
}
