package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"warehouse/application"
	"warehouse/domain/repository"
	"warehouse/infraestructure/serializer"
)

type InventoryController struct {
	inventoryApplication *application.InventoryApplication
}

func NewInventoryController(inventoryRepository repository.IInventoryRepository) *InventoryController {
	inventoryApplication := application.NewInventoryApplication(inventoryRepository)
	return &InventoryController{inventoryApplication}
}

func (inventoryController *InventoryController) Index(c *gin.Context) {
	inventories, error := inventoryController.inventoryApplication.All()
	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		return
	}
	c.JSON(http.StatusOK, serializer.NewInventoryListSerializerFromDomains(inventories))
}

func (inventoryController *InventoryController) Get(c *gin.Context) {
	id := c.Param("id")
	inventory, error := inventoryController.inventoryApplication.Show(id)
	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		return
	}

	if inventory == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "The requested resource doesn't exist (Inventory = " + id + ")"})
		return
	}

	c.JSON(http.StatusOK, serializer.NewInventorySerializerFromDomain(inventory))
}

func (inventoryController *InventoryController) Create(c *gin.Context) {
	var inventorySerializer serializer.InventorySerializer
	if err := c.ShouldBindJSON(&inventorySerializer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}
	inventory, error := inventoryController.inventoryApplication.Create(time.Time(inventorySerializer.OperationDate))
	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "error": error.Error() })
		return
	}
	c.JSON(http.StatusOK, serializer.NewInventorySerializerFromDomain(inventory))
}

func (inventoryController *InventoryController) Update(c *gin.Context) {
	id := c.Param("id")
	var inventorySerializer serializer.InventorySerializer
	if err := c.ShouldBindJSON(&inventorySerializer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}
	err := inventoryController.inventoryApplication.Update(id, time.Time(inventorySerializer.OperationDate))
	if err != nil {
		c.JSON(err.HttpStatusCode(), err.Error())
		return
	}
	inventorySerializer.Id = id
	c.JSON(http.StatusOK, inventorySerializer)
}