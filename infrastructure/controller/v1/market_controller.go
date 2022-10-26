package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"warehouse/application"
	"warehouse/infrastructure/controller/context"
	"warehouse/infrastructure/dto"
)

type MarketController struct {
	marketApplication *application.MarketApplication
}

func NewMarketController(marketApplication *application.MarketApplication) *MarketController {
	return &MarketController{marketApplication}
}

func (marketController *MarketController) Index(c context.IContextAdapter) {
	inventories, err := marketController.marketApplication.All()
	if err != nil {
		c.JSON(err.HttpStatusCode(), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.NewMarketListDtoFromDomains(inventories))
}

func (marketController *MarketController) Get(c context.IContextAdapter) {
	var marketDto dto.MarketDto
	if err := c.ShouldBindUri(&marketDto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	market, err := marketController.marketApplication.Show(marketDto.Id)
	if err != nil {
		c.JSON(err.HttpStatusCode(), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.NewMarketDtoFromDomain(market))
}

func (marketController *MarketController) Create(c context.IContextAdapter) {
	var marketDto dto.MarketDto
	if err := c.ShouldBindJSON(&marketDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	market, err := marketController.marketApplication.Create(marketDto.ToDomain())
	if err != nil {
		c.JSON(err.HttpStatusCode(), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, dto.NewMarketDtoFromDomain(market))
}

func (marketController *MarketController) Update(c context.IContextAdapter) {
	var marketDto dto.MarketDto
	if err := c.ShouldBindUri(&marketDto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&marketDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	market, err := marketController.marketApplication.Update(marketDto.ToDomain())
	if err != nil {
		c.JSON(err.HttpStatusCode(), err.Error())
		return
	}
	c.JSON(http.StatusOK, dto.NewMarketDtoFromDomain(market))
}

func (marketController *MarketController) Delete(c context.IContextAdapter) {
	var marketDto dto.MarketDto
	if err := c.ShouldBindUri(&marketDto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err := marketController.marketApplication.Delete(marketDto.ToDomain())
	if err != nil {
		c.JSON(err.HttpStatusCode(), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
