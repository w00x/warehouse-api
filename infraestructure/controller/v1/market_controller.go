package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"warehouse/application"
	"warehouse/domain/repository"
	"warehouse/infraestructure/serializer"
)

type MarketController struct {
	marketApplication *application.MarketApplication
}

func NewMarketController(marketRepository repository.IMarketRepository) *MarketController {
	marketApplication := application.NewMarketApplication(marketRepository)
	return &MarketController{marketApplication}
}

func (marketController *MarketController) Index(c *gin.Context) {
	inventories, err := marketController.marketApplication.All()
	if err != nil {
		c.JSON(err.HttpStatusCode(), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, serializer.NewMarketListSerializerFromDomains(inventories))
}

func (marketController *MarketController) Get(c *gin.Context) {
	var marketSerializer serializer.MarketSerializer
	if err := c.ShouldBindUri(&marketSerializer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	market, err := marketController.marketApplication.Show(marketSerializer.Id)
	if err != nil {
		c.JSON(err.HttpStatusCode(), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, serializer.NewMarketSerializerFromDomain(market))
}

func (marketController *MarketController) Create(c *gin.Context) {
	var marketSerializer serializer.MarketSerializer
	if err := c.ShouldBindJSON(&marketSerializer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}
	market, err := marketController.marketApplication.Create(marketSerializer.ToDomain())
	if err != nil {
		c.JSON(err.HttpStatusCode(), gin.H{ "error": err.Error() })
		return
	}
	c.JSON(http.StatusCreated, serializer.NewMarketSerializerFromDomain(market))
}

func (marketController *MarketController) Update(c *gin.Context) {
	var marketSerializer serializer.MarketSerializer
	if err := c.ShouldBindUri(&marketSerializer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&marketSerializer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}
	market, err := marketController.marketApplication.Update(marketSerializer.ToDomain())
	if err != nil {
		c.JSON(err.HttpStatusCode(), err.Error())
		return
	}
	c.JSON(http.StatusOK, serializer.NewMarketSerializerFromDomain(market))
}

func (marketController *MarketController) Delete(c *gin.Context) {
	var marketSerializer serializer.MarketSerializer
	if err := c.ShouldBindUri(&marketSerializer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err := marketController.marketApplication.Delete(marketSerializer.ToDomain())
	if err != nil {
		c.JSON(err.HttpStatusCode(), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}