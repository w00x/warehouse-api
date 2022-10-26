package dto

import (
	"github.com/stretchr/testify/assert"
	"testing"
	dom "warehouse/domain"
	"warehouse/infrastructure/dto"
	"warehouse/tests/factories"
)

func TestStockDto_ToDomain(t *testing.T) {
	domain := factories.NewStockFactory()
	stockDto := dto.NewStockResponseDtoFromDomain(domain)

	assert.Equal(t, stockDto.ToDomain(), domain)
}

func TestNewStockDto(t *testing.T) {
	domain := factories.NewStockFactory()
	stockDto := dto.NewStockResponseDtoFromDomain(domain)
	newStockDto := dto.StockResponseDto{
		Id:             domain.Id(),
		Item:           dto.NewItemDtoFromDomain(domain.Item),
		Rack:           dto.NewRackDtoFromDomain(domain.Rack),
		Quantity:       domain.Quantity,
		OperationDate:  domain.OperationDate,
		ExpirationDate: domain.ExpirationDate,
	}

	assert.Equal(t, stockDto.Id, newStockDto.Id)
	assert.Equal(t, stockDto.Quantity, newStockDto.Quantity)
}

func TestNewStockDtoFromDomain(t *testing.T) {
	domain := factories.NewStockFactory()
	stockDto := dto.NewStockResponseDtoFromDomain(domain)

	assert.Equal(t, stockDto.ToDomain(), domain)
}

func TestNewStockListDtoFromDomains(t *testing.T) {
	domain := factories.NewStockFactory()
	stocks := []dom.Stock{*domain}
	stockDtos := dto.NewStockListResponseDtoFromDomains(&stocks)

	assert.Equal(t, stockDtos[0].ToDomain(), domain)
}
