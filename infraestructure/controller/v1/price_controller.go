package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"warehouse/application"
	"warehouse/domain/repository"
	"warehouse/infraestructure/serializer"
)

type PriceController struct {
	priceApplication *application.PriceApplication
}

func NewPriceController(priceRepository repository.IPriceRepository) *PriceController {
	priceApplication := application.NewPriceApplication(priceRepository)
	return &PriceController{priceApplication}
}

func (priceController *PriceController) Index(c *gin.Context) {
	inventories, err := priceController.priceApplication.All()
	if err != nil {
		c.JSON(err.HttpStatusCode(), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, serializer.NewPriceResponseListSerializerFromDomains(inventories))
}

func (priceController *PriceController) Get(c *gin.Context) {
	var priceSerializer serializer.PriceSerializer
	if err := c.ShouldBindUri(&priceSerializer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	price, err := priceController.priceApplication.Show(priceSerializer.Id)
	if err != nil {
		c.JSON(err.HttpStatusCode(), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, serializer.NewPriceResponseSerializerFromDomain(price))
}

func (priceController *PriceController) Create(c *gin.Context) {
	var priceSerializer serializer.PriceSerializer
	if err := c.ShouldBindJSON(&priceSerializer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}
	price, err := priceController.priceApplication.Create(priceSerializer.ToDomain())
	if err != nil {
		c.JSON(err.HttpStatusCode(), gin.H{ "error": err.Error() })
		return
	}
	c.JSON(http.StatusOK, serializer.NewPriceResponseSerializerFromDomain(price))
}

func (priceController *PriceController) Update(c *gin.Context) {
	var priceSerializer serializer.PriceSerializer
	if err := c.ShouldBindUri(&priceSerializer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&priceSerializer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}
	price, err := priceController.priceApplication.Update(priceSerializer.ToDomain())
	if err != nil {
		c.JSON(err.HttpStatusCode(), err.Error())
		return
	}
	c.JSON(http.StatusOK, serializer.NewPriceResponseSerializerFromDomain(price))
}

func (priceController *PriceController) Delete(c *gin.Context) {
	var priceSerializer serializer.PriceSerializer
	if err := c.ShouldBindUri(&priceSerializer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err := priceController.priceApplication.Delete(priceSerializer.ToDomain())
	if err != nil {
		c.JSON(err.HttpStatusCode(), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}