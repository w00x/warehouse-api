package factories

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"warehouse/domain"
	"warehouse/infrastructure/repository/gorm"
)

type Item struct {
	Name                 string
	UnitSizePresentation string
	SizePresentation     int `fake:"{int16}"`
	Code                 string
	Container            string
	Photo                string
}

func (i Item) ToDomain() *domain.Item {
	return domain.NewItem("", i.Name, i.UnitSizePresentation, i.SizePresentation, i.Code, i.Container, i.Photo)
}

func FromItemDomainToFactory(item *domain.Item) *Item {
	return &Item{
		Name:                 item.Name,
		UnitSizePresentation: item.UnitSizePresentation,
		SizePresentation:     item.SizePresentation,
		Code:                 item.Code,
		Container:            item.Container,
		Photo:                item.Photo,
	}
}

func NewItemFactory() *domain.Item {
	item := &Item{}
	err := gofakeit.Struct(item)
	if err != nil {
		fmt.Println(err)
	}

	repo := gorm.NewItemRepository()
	itemDomain, errRepo := repo.Create(item.ToDomain())
	if errRepo != nil {
		panic(err)
	}

	return itemDomain
}

func NewItemDomainFactory() *domain.Item {
	item := &Item{}
	err := gofakeit.Struct(item)
	if err != nil {
		fmt.Println(err)
	}

	return item.ToDomain()
}

func NewItemObjectFactory() map[string]interface{} {
	item := &Item{}
	err := gofakeit.Struct(item)
	if err != nil {
		fmt.Println(err)
	}

	itemMarshal := map[string]interface{}{
		"name":                   item.Name,
		"unit_size_presentation": item.UnitSizePresentation,
		"size_presentation":      item.SizePresentation,
		"code":                   item.Code,
		"container":              item.Container,
		"photo":                  item.Photo,
	}

	return itemMarshal
}

func NewItemFactoryList(count int) []*domain.Item {
	var itemDomains []*domain.Item
	repo := gorm.NewItemRepository()

	for i := 0; i < count; i++ {
		item := &Item{}
		err := gofakeit.Struct(item)
		if err != nil {
			panic(err)
		}

		itemDomain, errRepo := repo.Create(item.ToDomain())
		if errRepo != nil {
			fmt.Println("Error: %w", errRepo)
		}
		itemDomains = append(itemDomains, itemDomain)
	}

	return itemDomains
}
