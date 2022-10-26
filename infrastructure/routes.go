package infrastructure

import (
	"github.com/gin-gonic/gin"
	"warehouse/infrastructure/factory"
)

func Routes(factoryAdapter string) *gin.Engine {
	route := gin.Default()
	contextAdapter := "gin"
	version1 := route.Group("/v1")
	{
		inventory := version1.Group("inventory")
		{
			inventoryController := InitializeInventoryController(factoryAdapter)
			inventory.GET("", func(context *gin.Context) {
				ctx, err := factory.ContextFactory(contextAdapter, context)
				if err != nil {
					panic(err)
				}
				inventoryController.Index(ctx)
			})
			inventory.GET("/:id", func(context *gin.Context) {
				ctx, err := factory.ContextFactory(contextAdapter, context)
				if err != nil {
					panic(err)
				}
				inventoryController.Get(ctx)
			})
			inventory.POST("", func(context *gin.Context) {
				ctx, err := factory.ContextFactory(contextAdapter, context)
				if err != nil {
					panic(err)
				}
				inventoryController.Create(ctx)
			})
			inventory.PATCH("/:id", func(context *gin.Context) {
				ctx, err := factory.ContextFactory(contextAdapter, context)
				if err != nil {
					panic(err)
				}
				inventoryController.Update(ctx)
			})
			inventory.DELETE("/:id", func(context *gin.Context) {
				ctx, err := factory.ContextFactory(contextAdapter, context)
				if err != nil {
					panic(err)
				}
				inventoryController.Delete(ctx)
			})
		}

		item := version1.Group("item")
		{
			itemController := InitializeItemController(factoryAdapter)
			item.GET("", func(context *gin.Context) {
				ctx, err := factory.ContextFactory(contextAdapter, context)
				if err != nil {
					panic(err)
				}
				itemController.Index(ctx)
			})
			item.GET("/:id", func(context *gin.Context) {
				ctx, err := factory.ContextFactory(contextAdapter, context)
				if err != nil {
					panic(err)
				}
				itemController.Get(ctx)
			})
			item.POST("", func(context *gin.Context) {
				ctx, err := factory.ContextFactory(contextAdapter, context)
				if err != nil {
					panic(err)
				}
				itemController.Create(ctx)
			})
			item.PATCH("/:id", func(context *gin.Context) {
				ctx, err := factory.ContextFactory(contextAdapter, context)
				if err != nil {
					panic(err)
				}
				itemController.Update(ctx)
			})
			item.DELETE("/:id", func(context *gin.Context) {
				ctx, err := factory.ContextFactory(contextAdapter, context)
				if err != nil {
					panic(err)
				}
				itemController.Delete(ctx)
			})
		}

		market := version1.Group("market")
		{
			marketController := InitializeMarketController(factoryAdapter)
			market.GET("", func(context *gin.Context) {
				ctx, err := factory.ContextFactory(contextAdapter, context)
				if err != nil {
					panic(err)
				}
				marketController.Index(ctx)
			})
			market.GET("/:id", func(context *gin.Context) {
				ctx, err := factory.ContextFactory(contextAdapter, context)
				if err != nil {
					panic(err)
				}
				marketController.Get(ctx)
			})
			market.POST("", func(context *gin.Context) {
				ctx, err := factory.ContextFactory(contextAdapter, context)
				if err != nil {
					panic(err)
				}
				marketController.Create(ctx)
			})
			market.PATCH("/:id", func(context *gin.Context) {
				ctx, err := factory.ContextFactory(contextAdapter, context)
				if err != nil {
					panic(err)
				}
				marketController.Update(ctx)
			})
			market.DELETE("/:id", func(context *gin.Context) {
				ctx, err := factory.ContextFactory(contextAdapter, context)
				if err != nil {
					panic(err)
				}
				marketController.Delete(ctx)
			})
		}

		price := version1.Group("price")
		{
			priceController := InitializePriceController(factoryAdapter)
			price.GET("", func(context *gin.Context) {
				ctx, err := factory.ContextFactory(contextAdapter, context)
				if err != nil {
					panic(err)
				}
				priceController.Index(ctx)
			})
			price.GET("/:id", func(context *gin.Context) {
				ctx, err := factory.ContextFactory(contextAdapter, context)
				if err != nil {
					panic(err)
				}
				priceController.Get(ctx)
			})
			price.POST("", func(context *gin.Context) {
				ctx, err := factory.ContextFactory(contextAdapter, context)
				if err != nil {
					panic(err)
				}
				priceController.Create(ctx)
			})
			price.PATCH("/:id", func(context *gin.Context) {
				ctx, err := factory.ContextFactory(contextAdapter, context)
				if err != nil {
					panic(err)
				}
				priceController.Update(ctx)
			})
			price.DELETE("/:id", func(context *gin.Context) {
				ctx, err := factory.ContextFactory(contextAdapter, context)
				if err != nil {
					panic(err)
				}
				priceController.Delete(ctx)
			})
		}

		rack := version1.Group("rack")
		{
			rackController := InitializeRackController(factoryAdapter)
			rack.GET("", func(context *gin.Context) {
				ctx, err := factory.ContextFactory(contextAdapter, context)
				if err != nil {
					panic(err)
				}
				rackController.Index(ctx)
			})
			rack.GET("/:id", func(context *gin.Context) {
				ctx, err := factory.ContextFactory(contextAdapter, context)
				if err != nil {
					panic(err)
				}
				rackController.Get(ctx)
			})
			rack.POST("", func(context *gin.Context) {
				ctx, err := factory.ContextFactory(contextAdapter, context)
				if err != nil {
					panic(err)
				}
				rackController.Create(ctx)
			})
			rack.PATCH("/:id", func(context *gin.Context) {
				ctx, err := factory.ContextFactory(contextAdapter, context)
				if err != nil {
					panic(err)
				}
				rackController.Update(ctx)
			})
			rack.DELETE("/:id", func(context *gin.Context) {
				ctx, err := factory.ContextFactory(contextAdapter, context)
				if err != nil {
					panic(err)
				}
				rackController.Delete(ctx)
			})
		}

		stock := version1.Group("stock")
		{
			stockController := InitializeStockController(factoryAdapter)
			stock.GET("", func(context *gin.Context) {
				ctx, err := factory.ContextFactory(contextAdapter, context)
				if err != nil {
					panic(err)
				}
				stockController.Index(ctx)
			})

			stock.GET("/:id", func(context *gin.Context) {
				ctx, err := factory.ContextFactory(contextAdapter, context)
				if err != nil {
					panic(err)
				}
				stockController.Get(ctx)
			})

			stock.POST("", func(context *gin.Context) {
				ctx, err := factory.ContextFactory(contextAdapter, context)
				if err != nil {
					panic(err)
				}
				stockController.Create(ctx)
			})
		}
	}

	return route
}
