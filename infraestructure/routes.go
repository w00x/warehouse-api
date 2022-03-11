package infraestructure

import (
	"github.com/gin-gonic/gin"
	"warehouse/infraestructure/controller/v1"
	"warehouse/infraestructure/repository/postgres"
)

func Routes() {
	route := gin.Default()
	version1 := route.Group("/v1")
	{
		inventory := version1.Group("inventory")
		{
			inventoryController := v1.NewInventoryController(GetInventoryRepository())
			inventory.GET("", inventoryController.Index)
			inventory.GET("/:id", inventoryController.Get)
			inventory.POST("", inventoryController.Create)
			inventory.PATCH("/:id", inventoryController.Update)
			inventory.DELETE("/:id", inventoryController.Delete)
		}

		item := version1.Group("item")
		{
			itemController := v1.NewItemController(GetItemRepository())
			item.GET("", itemController.Index)
			item.GET("/:id", itemController.Get)
			item.POST("", itemController.Create)
			item.PATCH("/:id", itemController.Update)
			item.DELETE("/:id", itemController.Delete)
		}

		market := version1.Group("market")
		{
			marketController := v1.NewMarketController(GetMarketRepository())
			market.GET("", marketController.Index)
			market.GET("/:id", marketController.Get)
			market.POST("", marketController.Create)
			market.PATCH("/:id", marketController.Update)
			market.DELETE("/:id", marketController.Delete)
		}

		price := version1.Group("price")
		{
			priceController := v1.NewPriceController(GetPriceRepository())
			price.GET("", priceController.Index)
			price.GET("/:id", priceController.Get)
			price.POST("", priceController.Create)
			price.PATCH("/:id", priceController.Update)
			price.DELETE("/:id", priceController.Delete)
		}

		rack := version1.Group("rack")
		{
			rackController := v1.NewRackController(GetRackRepository())
			rack.GET("", rackController.Index)
			rack.GET("/:id", rackController.Get)
			rack.POST("", rackController.Create)
			rack.PATCH("/:id", rackController.Update)
			rack.DELETE("/:id", rackController.Delete)
		}

		stock := version1.Group("stock")
		{
			stockController := v1.NewStockController(GetStockRepository())
			stock.GET("", stockController.Index)
			stock.GET("/:id", stockController.Get)
			stock.POST("", stockController.Create)
			stock.PATCH("/:id", stockController.Update)
			stock.DELETE("/:id", stockController.Delete)
		}
	}
	route.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func GetInventoryRepository() *postgres.InventoryRepository {
	return postgres.NewInventoryRepository()
}

func GetItemRepository() *postgres.ItemRepository {
	return postgres.NewItemRepository()
}

func GetMarketRepository() *postgres.MarketRepository {
	return postgres.NewMarketRepository()
}

func GetPriceRepository() *postgres.PriceRepository {
	return postgres.NewPriceRepository()
}

func GetRackRepository() *postgres.RackRepository {
	return postgres.NewRackRepository()
}

func GetStockRepository() *postgres.StockRepository {
	return postgres.NewStockRepository()
}