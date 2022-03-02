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
	}
	route.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func GetInventoryRepository() *postgres.InventoryRepository {
	return postgres.NewInventoryRepository()
}
