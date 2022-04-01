package dto

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"testing"
	dom "warehouse/domain"
	"warehouse/infraestructure/dto"
	"warehouse/shared"
)

func TestInventoryDto_ToDomain(t *testing.T) {
	id := uint(gofakeit.Number(1,10))
	opertionDate := shared.DateTime(gofakeit.Date())
	inventoryDto := dto.NewInventoryDto(id, opertionDate)
	domain := dom.NewInventory(id, opertionDate)

	assert.Equal(t, inventoryDto.ToDomain(), domain)
}

func TestNewInventoryDto(t *testing.T) {
	id := uint(gofakeit.Number(1,10))
	opertionDate := shared.DateTime(gofakeit.Date())
	inventoryDto := dto.NewInventoryDto(id, opertionDate)
	newInventoryDto := dto.InventoryDto{
		Id:            id,
		OperationDate: opertionDate,
	}

	assert.Equal(t, inventoryDto.Id, newInventoryDto.Id)
	assert.Equal(t, inventoryDto.OperationDate, newInventoryDto.OperationDate)
}

func TestNewInventoryDtoFromDomain(t *testing.T) {
	id := uint(gofakeit.Number(1,10))
	opertionDate := shared.DateTime(gofakeit.Date())
	domain := dom.NewInventory(id, opertionDate)
	inventoryDto := dto.NewInventoryDtoFromDomain(domain)

	assert.Equal(t, inventoryDto.ToDomain(), domain)
}

func TestNewInventoryListDtoFromDomains(t *testing.T) {
	id := uint(gofakeit.Number(1,10))
	opertionDate := shared.DateTime(gofakeit.Date())
	domain := dom.NewInventory(id, opertionDate)
	inventories := []dom.Inventory { *domain }
	inventoryDtos := dto.NewInventoryListDtoFromDomains(&inventories)

	assert.Equal(t, inventoryDtos[0].ToDomain(), domain)
}
