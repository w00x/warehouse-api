package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"warehouse/application"
	"warehouse/domain/repository"
	"warehouse/infraestructure/serializer"
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
	c.JSON(http.StatusOK, serializer.NewStockListSerializerFromDomains(inventories))
}

func (stockController *StockController) Get(c *gin.Context) {
	var stockSerializer serializer.StockSerializer
	if err := c.ShouldBindUri(&stockSerializer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	stock, err := stockController.stockApplication.Show(stockSerializer.Id)
	if err != nil {
		c.JSON(err.HttpStatusCode(), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, serializer.NewStockSerializerFromDomain(stock))
}

func (stockController *StockController) Create(c *gin.Context) {
	var stockSerializer serializer.StockSerializer
	if err := c.ShouldBindJSON(&stockSerializer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}
	stock, err := stockController.stockApplication.Create(stockSerializer.ToDomain())
	if err != nil {
		c.JSON(err.HttpStatusCode(), gin.H{ "error": err.Error() })
		return
	}
	c.JSON(http.StatusOK, serializer.NewStockSerializerFromDomain(stock))
}

func (stockController *StockController) Update(c *gin.Context) {
	var stockSerializer serializer.StockSerializer
	if err := c.ShouldBindUri(&stockSerializer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&stockSerializer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}
	stock, err := stockController.stockApplication.Update(stockSerializer.ToDomain())
	if err != nil {
		c.JSON(err.HttpStatusCode(), err.Error())
		return
	}
	c.JSON(http.StatusOK, serializer.NewStockSerializerFromDomain(stock))
}

func (stockController *StockController) Delete(c *gin.Context) {
	var stockSerializer serializer.StockSerializer
	if err := c.ShouldBindUri(&stockSerializer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err := stockController.stockApplication.Delete(stockSerializer.ToDomain())
	if err != nil {
		c.JSON(err.HttpStatusCode(), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}