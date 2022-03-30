package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"warehouse/application"
	"warehouse/domain/repository"
	"warehouse/infraestructure/dto"
)

type ItemController struct {
	itemApplication *application.ItemApplication
}

func NewItemController(itemRepository repository.IItemRepository) *ItemController {
	itemApplication := application.NewItemApplication(itemRepository)
	return &ItemController{itemApplication}
}

func (itemController *ItemController) Index(c *gin.Context) {
	inventories, err := itemController.itemApplication.All()
	if err != nil {
		c.JSON(err.HttpStatusCode(), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.NewItemListDtoFromDomains(inventories))
}

func (itemController *ItemController) Get(c *gin.Context) {
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

func (itemController *ItemController) Create(c *gin.Context) {
	var itemDto dto.ItemDto
	if err := c.ShouldBindJSON(&itemDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}
	item, err := itemController.itemApplication.Create(itemDto.ToDomain())
	if err != nil {
		c.JSON(err.HttpStatusCode(), gin.H{ "error": err.Error() })
		return
	}
	c.JSON(http.StatusCreated, dto.NewItemDtoFromDomain(item))
}

func (itemController *ItemController) Update(c *gin.Context) {
	var itemDto dto.ItemDto
	if err := c.ShouldBindUri(&itemDto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&itemDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}
	item, err := itemController.itemApplication.Update(itemDto.ToDomain())
	if err != nil {
		c.JSON(err.HttpStatusCode(), err.Error())
		return
	}
	c.JSON(http.StatusOK, dto.NewItemDtoFromDomain(item))
}

func (itemController *ItemController) Delete(c *gin.Context) {
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