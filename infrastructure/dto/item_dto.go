package dto

import "warehouse/domain"

type ItemDto struct {
	Id                   string `json:"id" uri:"id"`
	Name                 string `json:"name"`
	UnitSizePresentation string `json:"unit_size_presentation"`
	SizePresentation     int    `json:"size_presentation"`
	Code                 string `json:"code"`
	Container            string `json:"container"`
	Photo                string `json:"photo"`
}

func NewItemDto(id string, name string, unitSizePresentation string, sizePresentation int,
	code string, container string, photo string) *ItemDto {
	return &ItemDto{Id: id, Name: name, UnitSizePresentation: unitSizePresentation,
		SizePresentation: sizePresentation, Code: code, Container: container, Photo: photo}
}

func NewItemDtoFromDomain(item *domain.Item) *ItemDto {
	if item == nil {
		return nil
	}
	return NewItemDto(item.Id(), item.Name, item.UnitSizePresentation, item.SizePresentation,
		item.Code, item.Container, item.Photo)
}

func NewItemListDtoFromDomains(items *[]domain.Item) []*ItemDto {
	var itemDtos []*ItemDto
	for _, item := range *items {
		itemDtos = append(itemDtos,
			NewItemDtoFromDomain(&item))
	}
	return itemDtos
}

func (ise ItemDto) ToDomain() *domain.Item {
	return domain.NewItem(ise.Id, ise.Name, ise.UnitSizePresentation, ise.SizePresentation,
		ise.Code, ise.Container, ise.Photo)
}
