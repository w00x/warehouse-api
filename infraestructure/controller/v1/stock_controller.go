package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"warehouse/application"
	"warehouse/domain/repository"
	"warehouse/infraestructure/dto"
)

type StockController struct {
	stockApplication *application.StockApplication
}

func NewStockController(stockRepository repository.IStockRepository) *StockController {
	stockApplication := application.NewStockApplication(stockRepository)
	return &StockController{stockApplication}
}

func (stockController *StockController) Index(c *gin.Context) {
	inventories, err := stockController.stockApplication.All()
	if err != nil {
		c.JSON(err.HttpStatusCode(), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.NewStockListResponseDtoFromDomains(inventories))
}

func (stockController *StockController) Get(c *gin.Context) {
	var stockDto dto.StockRequestDto
	if err := c.ShouldBindUri(&stockDto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	stock, err := stockController.stockApplication.Show(stockDto.Id)
	if err != nil {
		c.JSON(err.HttpStatusCode(), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.NewStockResponseDtoFromDomain(stock))
}

func (stockController *StockController) Create(c *gin.Context) {
	var stockDto dto.StockRequestDto
	if err := c.ShouldBindJSON(&stockDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}
	stock, err := stockController.stockApplication.Create(stockDto.ToDomain())
	if err != nil {
		c.JSON(err.HttpStatusCode(), gin.H{ "error": err.Error() })
		return
	}
	c.JSON(http.StatusCreated, dto.NewStockResponseDtoFromDomain(stock))
}

func (stockController *StockController) Update(c *gin.Context) {
	var stockDto dto.StockRequestDto
	if err := c.ShouldBindUri(&stockDto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&stockDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}
	stock, err := stockController.stockApplication.Update(stockDto.ToDomain())
	if err != nil {
		c.JSON(err.HttpStatusCode(), err.Error())
		return
	}
	c.JSON(http.StatusOK, dto.NewStockResponseDtoFromDomain(stock))
}

func (stockController *StockController) Delete(c *gin.Context) {
	var stockDto dto.StockRequestDto
	if err := c.ShouldBindUri(&stockDto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err := stockController.stockApplication.Delete(stockDto.ToDomain())
	if err != nil {
		c.JSON(err.HttpStatusCode(), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}