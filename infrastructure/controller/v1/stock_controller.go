package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"warehouse/application"
	"warehouse/infrastructure/controller/context"
	"warehouse/infrastructure/dto"
)

type StockController struct {
	stockApplication *application.StockApplication
}

func NewStockController(stockApplication *application.StockApplication) *StockController {
	return &StockController{stockApplication}
}

func (stockController *StockController) Index(c context.IContextAdapter) {
	stocks, err := stockController.stockApplication.All()
	if err != nil {
		c.JSON(err.HttpStatusCode(), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.NewStockListResponseDtoFromDomains(stocks))
}

func (stockController *StockController) Create(c context.IContextAdapter) {
	var stockDto dto.StockRequestDto
	if err := c.ShouldBindJSON(&stockDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	stock, err := stockController.stockApplication.Create(stockDto.ToDomain())
	if err != nil {
		c.JSON(err.HttpStatusCode(), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, dto.NewStockResponseDtoFromDomain(stock))
}

func (stockController *StockController) Get(c context.IContextAdapter) {
	var itemDto dto.ItemDto
	if err := c.ShouldBindUri(&itemDto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	stock, err := stockController.stockApplication.Show(itemDto.Id)

	if err != nil {
		c.JSON(err.HttpStatusCode(), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.NewStockResponseDtoFromDomain(stock))
}

func (stockController StockController) AllByInventory(c context.IContextAdapter) {
	var inventoryDto dto.ItemDto
	if err := c.ShouldBindUri(&inventoryDto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	stocks, err := stockController.stockApplication.AllByInventory(inventoryDto.Id)

	if err != nil {
		c.JSON(err.HttpStatusCode(), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.NewStockListResponseDtoFromDomains(stocks))
}
