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
	inventories, err := inventoryController.inventoryApplication.All()
	if err != nil {
		c.JSON(err.HttpStatusCode(), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, serializer.NewInventoryListSerializerFromDomains(inventories))
}

func (inventoryController *InventoryController) Get(c *gin.Context) {
	var inventorySerializer serializer.InventorySerializer
	if err := c.ShouldBindUri(&inventorySerializer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	inventory, err := inventoryController.inventoryApplication.Show(inventorySerializer.Id)
	if err != nil {
		c.JSON(err.HttpStatusCode(), gin.H{"error": err.Error()})
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
	inventory, err := inventoryController.inventoryApplication.Create(time.Time(inventorySerializer.OperationDate))
	if err != nil {
		c.JSON(err.HttpStatusCode(), gin.H{ "error": err.Error() })
		return
	}
	c.JSON(http.StatusOK, serializer.NewInventorySerializerFromDomain(inventory))
}

func (inventoryController *InventoryController) Update(c *gin.Context) {
	var inventorySerializer serializer.InventorySerializer
	if err := c.ShouldBindUri(&inventorySerializer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&inventorySerializer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}
	inventory, err := inventoryController.inventoryApplication.Update(inventorySerializer.Id, time.Time(inventorySerializer.OperationDate))
	if err != nil {
		c.JSON(err.HttpStatusCode(), err.Error())
		return
	}
	c.JSON(http.StatusOK, serializer.NewInventorySerializerFromDomain(inventory))
}

func (inventoryController *InventoryController) Delete(c *gin.Context) {
	var inventorySerializer serializer.InventorySerializer
	if err := c.ShouldBindUri(&inventorySerializer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err := inventoryController.inventoryApplication.Delete(inventorySerializer.Id)
	if err != nil {
		c.JSON(err.HttpStatusCode(), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}