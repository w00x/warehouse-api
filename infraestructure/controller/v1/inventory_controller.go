package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"warehouse/application"
	"warehouse/domain/repository"
	"warehouse/infraestructure/dto"
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
	c.JSON(http.StatusOK, dto.NewInventoryListDtoFromDomains(inventories))
}

func (inventoryController *InventoryController) Get(c *gin.Context) {
	var inventoryDto dto.InventoryDto
	if err := c.ShouldBindUri(&inventoryDto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	inventory, err := inventoryController.inventoryApplication.Show(inventoryDto.Id)
	if err != nil {
		c.JSON(err.HttpStatusCode(), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.NewInventoryDtoFromDomain(inventory))
}

func (inventoryController *InventoryController) Create(c *gin.Context) {
	var inventoryDto dto.InventoryDto
	if err := c.ShouldBindJSON(&inventoryDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}
	inventory, err := inventoryController.inventoryApplication.Create(inventoryDto.ToDomain())
	if err != nil {
		c.JSON(err.HttpStatusCode(), gin.H{ "error": err.Error() })
		return
	}
	c.JSON(http.StatusCreated, dto.NewInventoryDtoFromDomain(inventory))
}

func (inventoryController *InventoryController) Update(c *gin.Context) {
	var inventoryDto dto.InventoryDto
	if err := c.ShouldBindUri(&inventoryDto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&inventoryDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}
	inventory, err := inventoryController.inventoryApplication.Update(inventoryDto.ToDomain())
	if err != nil {
		c.JSON(err.HttpStatusCode(), err.Error())
		return
	}
	c.JSON(http.StatusOK, dto.NewInventoryDtoFromDomain(inventory))
}

func (inventoryController *InventoryController) Delete(c *gin.Context) {
	var inventoryDto dto.InventoryDto
	if err := c.ShouldBindUri(&inventoryDto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err := inventoryController.inventoryApplication.Delete(inventoryDto.ToDomain())
	if err != nil {
		c.JSON(err.HttpStatusCode(), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}