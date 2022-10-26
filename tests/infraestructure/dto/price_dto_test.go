package dto

import (
	"github.com/stretchr/testify/assert"
	"testing"
	dom "warehouse/domain"
	"warehouse/infrastructure/dto"
	"warehouse/tests/factories"
)

func TestPriceDto_ToDomain(t *testing.T) {
	domain := factories.NewPriceFactory()
	priceDto := dto.NewPriceResponseDtoFromDomain(domain)

	assert.Equal(t, priceDto.ToDomain(), domain)
}

func TestNewPriceDto(t *testing.T) {
	domain := factories.NewPriceFactory()
	priceDto := dto.NewPriceResponseDtoFromDomain(domain)
	newPriceDto := dto.PriceResponseDto{
		Id:    domain.Id(),
		Item:  dto.NewItemDtoFromDomain(domain.Item),
		Price: domain.Price,
		Date:  domain.Date,
	}

	assert.Equal(t, priceDto.Id, newPriceDto.Id)
	assert.Equal(t, priceDto.Price, newPriceDto.Price)
}

func TestNewPriceDtoFromDomain(t *testing.T) {
	domain := factories.NewPriceFactory()
	priceDto := dto.NewPriceResponseDtoFromDomain(domain)

	assert.Equal(t, priceDto.ToDomain(), domain)
}

func TestNewPriceListDtoFromDomains(t *testing.T) {
	domain := factories.NewPriceFactory()
	prices := []dom.Price{*domain}
	priceDtos := dto.NewPriceResponseListDtoFromDomains(&prices)

	assert.Equal(t, priceDtos[0].ToDomain(), domain)
}
