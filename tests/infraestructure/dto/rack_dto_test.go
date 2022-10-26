package dto

import (
	"github.com/stretchr/testify/assert"
	"testing"
	dom "warehouse/domain"
	"warehouse/infrastructure/dto"
	"warehouse/tests/factories"
)

func TestRackDto_ToDomain(t *testing.T) {
	domain := factories.NewRackFactory()
	rackDto := dto.NewRackDtoFromDomain(domain)

	assert.Equal(t, rackDto.ToDomain(), domain)
}

func TestNewRackDto(t *testing.T) {
	domain := factories.NewRackFactory()
	rackDto := dto.NewRackDtoFromDomain(domain)
	newRackDto := dto.RackDto{
		Id:   domain.Id(),
		Name: domain.Name,
		Code: domain.Code,
	}

	assert.Equal(t, rackDto.Id, newRackDto.Id)
	assert.Equal(t, rackDto.Name, newRackDto.Name)
}

func TestNewRackDtoFromDomain(t *testing.T) {
	domain := factories.NewRackFactory()
	rackDto := dto.NewRackDtoFromDomain(domain)

	assert.Equal(t, rackDto.ToDomain(), domain)
}

func TestNewRackListDtoFromDomains(t *testing.T) {
	domain := factories.NewRackFactory()
	racks := []dom.Rack{*domain}
	rackDtos := dto.NewRackListDtoFromDomains(&racks)

	assert.Equal(t, rackDtos[0].ToDomain(), domain)
}
