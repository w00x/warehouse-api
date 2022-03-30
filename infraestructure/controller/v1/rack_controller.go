package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"warehouse/application"
	"warehouse/domain/repository"
	"warehouse/infraestructure/dto"
)

type RackController struct {
	rackApplication *application.RackApplication
}

func NewRackController(rackRepository repository.IRackRepository) *RackController {
	rackApplication := application.NewRackApplication(rackRepository)
	return &RackController{rackApplication}
}

func (rackController *RackController) Index(c *gin.Context) {
	inventories, err := rackController.rackApplication.All()
	if err != nil {
		c.JSON(err.HttpStatusCode(), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.NewRackListDtoFromDomains(inventories))
}

func (rackController *RackController) Get(c *gin.Context) {
	var rackDto dto.RackDto
	if err := c.ShouldBindUri(&rackDto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rack, err := rackController.rackApplication.Show(rackDto.Id)
	if err != nil {
		c.JSON(err.HttpStatusCode(), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.NewRackDtoFromDomain(rack))
}

func (rackController *RackController) Create(c *gin.Context) {
	var rackDto dto.RackDto
	if err := c.ShouldBindJSON(&rackDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}
	rack, err := rackController.rackApplication.Create(rackDto.ToDomain())
	if err != nil {
		c.JSON(err.HttpStatusCode(), gin.H{ "error": err.Error() })
		return
	}
	c.JSON(http.StatusCreated, dto.NewRackDtoFromDomain(rack))
}

func (rackController *RackController) Update(c *gin.Context) {
	var rackDto dto.RackDto
	if err := c.ShouldBindUri(&rackDto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&rackDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}
	rack, err := rackController.rackApplication.Update(rackDto.ToDomain())
	if err != nil {
		c.JSON(err.HttpStatusCode(), err.Error())
		return
	}
	c.JSON(http.StatusOK, dto.NewRackDtoFromDomain(rack))
}

func (rackController *RackController) Delete(c *gin.Context) {
	var rackDto dto.RackDto
	if err := c.ShouldBindUri(&rackDto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err := rackController.rackApplication.Delete(rackDto.ToDomain())
	if err != nil {
		c.JSON(err.HttpStatusCode(), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}