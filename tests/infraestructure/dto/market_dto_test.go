package dto

import (
	"github.com/stretchr/testify/assert"
	"testing"
	dom "warehouse/domain"
	"warehouse/infrastructure/dto"
	"warehouse/tests/factories"
)

func TestMarketDto_ToDomain(t *testing.T) {
	domain := factories.NewMarketFactory()
	marketDto := dto.NewMarketDtoFromDomain(domain)

	assert.Equal(t, marketDto.ToDomain(), domain)
}

func TestNewMarketDto(t *testing.T) {
	domain := factories.NewMarketFactory()
	marketDto := dto.NewMarketDtoFromDomain(domain)
	newMarketDto := dto.MarketDto{
		Id:   domain.Id(),
		Name: domain.Name,
	}

	assert.Equal(t, marketDto.Id, newMarketDto.Id)
	assert.Equal(t, marketDto.Name, newMarketDto.Name)
}

func TestNewMarketDtoFromDomain(t *testing.T) {
	domain := factories.NewMarketFactory()
	marketDto := dto.NewMarketDtoFromDomain(domain)

	assert.Equal(t, marketDto.ToDomain(), domain)
}

func TestNewMarketListDtoFromDomains(t *testing.T) {
	domain := factories.NewMarketFactory()
	markets := []dom.Market{*domain}
	marketDtos := dto.NewMarketListDtoFromDomains(&markets)

	assert.Equal(t, marketDtos[0].ToDomain(), domain)
}
