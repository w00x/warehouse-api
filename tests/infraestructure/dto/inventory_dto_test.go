package dto

import (
	"github.com/stretchr/testify/assert"
	"testing"
	dom "warehouse/domain"
	"warehouse/infrastructure/dto"
	"warehouse/tests/factories"
)

func TestInventoryDto_ToDomain(t *testing.T) {
	domain := factories.NewInventoryFactory()
	inventoryDto := dto.NewInventoryDtoFromDomain(domain)

	assert.Equal(t, inventoryDto.ToDomain(), domain)
}

func TestNewInventoryDto(t *testing.T) {
	domain := factories.NewInventoryFactory()
	inventoryDto := dto.NewInventoryDtoFromDomain(domain)
	newInventoryDto := dto.InventoryDto{
		Id:            domain.Id(),
		OperationDate: domain.OperationDate,
	}

	assert.Equal(t, inventoryDto.Id, newInventoryDto.Id)
	assert.Equal(t, inventoryDto.OperationDate, newInventoryDto.OperationDate)
}

func TestNewInventoryDtoFromDomain(t *testing.T) {
	domain := factories.NewInventoryFactory()
	inventoryDto := dto.NewInventoryDtoFromDomain(domain)

	assert.Equal(t, inventoryDto.ToDomain(), domain)
}

func TestNewInventoryListDtoFromDomains(t *testing.T) {
	domain := factories.NewInventoryFactory()
	inventories := []dom.Inventory{*domain}
	inventoryDtos := dto.NewInventoryListDtoFromDomains(&inventories)

	assert.Equal(t, inventoryDtos[0].ToDomain(), domain)
}
