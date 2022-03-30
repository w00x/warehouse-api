package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"warehouse/application"
	"warehouse/domain/repository"
	"warehouse/infraestructure/dto"
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
	c.JSON(http.StatusOK, dto.NewPriceResponseListDtoFromDomains(inventories))
}

func (priceController *PriceController) Get(c *gin.Context) {
	var priceDto dto.PriceRequesDto
	if err := c.ShouldBindUri(&priceDto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	price, err := priceController.priceApplication.Show(priceDto.Id)
	if err != nil {
		c.JSON(err.HttpStatusCode(), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.NewPriceResponseDtoFromDomain(price))
}

func (priceController *PriceController) Create(c *gin.Context) {
	var priceDto dto.PriceRequesDto
	if err := c.ShouldBindJSON(&priceDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}
	price, err := priceController.priceApplication.Create(priceDto.ToDomain())
	if err != nil {
		c.JSON(err.HttpStatusCode(), gin.H{ "error": err.Error() })
		return
	}
	c.JSON(http.StatusCreated, dto.NewPriceResponseDtoFromDomain(price))
}

func (priceController *PriceController) Update(c *gin.Context) {
	var priceDto dto.PriceRequesDto
	if err := c.ShouldBindUri(&priceDto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&priceDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}
	price, err := priceController.priceApplication.Update(priceDto.ToDomain())
	if err != nil {
		c.JSON(err.HttpStatusCode(), err.Error())
		return
	}
	c.JSON(http.StatusOK, dto.NewPriceResponseDtoFromDomain(price))
}

func (priceController *PriceController) Delete(c *gin.Context) {
	var priceDto dto.PriceRequesDto
	if err := c.ShouldBindUri(&priceDto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err := priceController.priceApplication.Delete(priceDto.ToDomain())
	if err != nil {
		c.JSON(err.HttpStatusCode(), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}