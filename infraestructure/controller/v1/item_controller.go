package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"warehouse/application"
	"warehouse/domain/repository"
	"warehouse/infraestructure/serializer"
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
	c.JSON(http.StatusOK, serializer.NewItemListSerializerFromDomains(inventories))
}

func (itemController *ItemController) Get(c *gin.Context) {
	var itemSerializer serializer.ItemSerializer
	if err := c.ShouldBindUri(&itemSerializer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	item, err := itemController.itemApplication.Show(itemSerializer.Id)
	if err != nil {
		c.JSON(err.HttpStatusCode(), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, serializer.NewItemSerializerFromDomain(item))
}

func (itemController *ItemController) Create(c *gin.Context) {
	var itemSerializer serializer.ItemSerializer
	if err := c.ShouldBindJSON(&itemSerializer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}
	item, err := itemController.itemApplication.Create(itemSerializer.ToDomain())
	if err != nil {
		c.JSON(err.HttpStatusCode(), gin.H{ "error": err.Error() })
		return
	}
	c.JSON(http.StatusOK, serializer.NewItemSerializerFromDomain(item))
}

func (itemController *ItemController) Update(c *gin.Context) {
	var itemSerializer serializer.ItemSerializer
	if err := c.ShouldBindUri(&itemSerializer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&itemSerializer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}
	item, err := itemController.itemApplication.Update(itemSerializer.ToDomain())
	if err != nil {
		c.JSON(err.HttpStatusCode(), err.Error())
		return
	}
	c.JSON(http.StatusOK, serializer.NewItemSerializerFromDomain(item))
}

func (itemController *ItemController) Delete(c *gin.Context) {
	var itemSerializer serializer.ItemSerializer
	if err := c.ShouldBindUri(&itemSerializer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err := itemController.itemApplication.Delete(itemSerializer.ToDomain())
	if err != nil {
		c.JSON(err.HttpStatusCode(), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}