package dto

import (
	"github.com/stretchr/testify/assert"
	"testing"
	dom "warehouse/domain"
	"warehouse/infrastructure/dto"
	"warehouse/tests/factories"
)

func TestItemDto_ToDomain(t *testing.T) {
	domain := factories.NewItemFactory()
	itemDto := dto.NewItemDtoFromDomain(domain)

	assert.Equal(t, itemDto.ToDomain(), domain)
}

func TestNewItemDto(t *testing.T) {
	domain := factories.NewItemFactory()
	itemDto := dto.NewItemDtoFromDomain(domain)
	newItemDto := dto.ItemDto{
		Id:                   domain.Id(),
		Name:                 domain.Name,
		UnitSizePresentation: domain.UnitSizePresentation,
		SizePresentation:     domain.SizePresentation,
		Code:                 domain.Code,
		Container:            domain.Container,
		Photo:                domain.Photo,
	}

	assert.Equal(t, itemDto.Id, newItemDto.Id)
	assert.Equal(t, itemDto.Name, newItemDto.Name)
}

func TestNewItemDtoFromDomain(t *testing.T) {
	domain := factories.NewItemFactory()
	itemDto := dto.NewItemDtoFromDomain(domain)

	assert.Equal(t, itemDto.ToDomain(), domain)
}

func TestNewItemListDtoFromDomains(t *testing.T) {
	domain := factories.NewItemFactory()
	items := []dom.Item{*domain}
	itemDtos := dto.NewItemListDtoFromDomains(&items)

	assert.Equal(t, itemDtos[0].ToDomain(), domain)
}
