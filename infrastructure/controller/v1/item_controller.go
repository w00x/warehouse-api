package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"warehouse/application"
	"warehouse/infrastructure/controller/context"
	"warehouse/infrastructure/dto"
)

type ItemController struct {
	itemApplication *application.ItemApplication
}

func NewItemController(itemApplication *application.ItemApplication) *ItemController {
	return &ItemController{itemApplication}
}

func (itemController *ItemController) Index(c context.IContextAdapter) {
	inventories, err := itemController.itemApplication.All()
	if err != nil {
		c.JSON(err.HttpStatusCode(), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.NewItemListDtoFromDomains(inventories))
}

func (itemController *ItemController) Get(c context.IContextAdapter) {
	var itemDto dto.ItemDto
	if err := c.ShouldBindUri(&itemDto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	item, err := itemController.itemApplication.Show(itemDto.Id)
	if err != nil {
		c.JSON(err.HttpStatusCode(), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.NewItemDtoFromDomain(item))
}

func (itemController *ItemController) Create(c context.IContextAdapter) {
	var itemDto dto.ItemDto
	if err := c.ShouldBindJSON(&itemDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	item, err := itemController.itemApplication.Create(itemDto.ToDomain())
	if err != nil {
		c.JSON(err.HttpStatusCode(), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, dto.NewItemDtoFromDomain(item))
}

func (itemController *ItemController) Update(c context.IContextAdapter) {
	var itemDto dto.ItemDto
	if err := c.ShouldBindUri(&itemDto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&itemDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	item, err := itemController.itemApplication.Update(itemDto.ToDomain())
	if err != nil {
		c.JSON(err.HttpStatusCode(), err.Error())
		return
	}
	c.JSON(http.StatusOK, dto.NewItemDtoFromDomain(item))
}

func (itemController *ItemController) Delete(c context.IContextAdapter) {
	var itemDto dto.ItemDto
	if err := c.ShouldBindUri(&itemDto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err := itemController.itemApplication.Delete(itemDto.ToDomain())
	if err != nil {
		c.JSON(err.HttpStatusCode(), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
