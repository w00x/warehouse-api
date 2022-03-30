package factories

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"testing"
	"warehouse/domain"
	"warehouse/infraestructure/repository/postgres"
)

type Item struct {
	Name 					string
	UnitSizePresentation 	string
	SizePresentation 		int
	Code 					string
	Container 				string
	Photo 					string
}

func (i Item) ToDomain() *domain.Item {
	return domain.NewItem(0, i.Name, i.UnitSizePresentation, i.SizePresentation, i.Code, i.Container, i.Photo)
}

func NewItemFactory(t *testing.T) *domain.Item {
	Item := &Item{}
	err := faker.FakeData(Item)
	if err != nil {
		fmt.Println(err)
	}

	repo := postgres.NewItemRepository()
	ItemDomain, errRepo := repo.Create(Item.ToDomain())
	if errRepo != nil {
		panic(err)
	}

	t.Cleanup(func() {
		CleanItem()
	})

	return ItemDomain
}

func NewItemObjectFactory() map[string]interface{} {
	item := &Item{}
	err := faker.FakeData(item)
	if err != nil {
		fmt.Println(err)
	}

	itemMarshal := map[string]interface{}{
		"name": item.Name,
		"unit_size_presentation": item.UnitSizePresentation,
		"size_presentation": item.SizePresentation,
		"code": item.Code,
		"container": item.Container,
		"photo": item.Photo,
	}

	return itemMarshal
}

func NewItemFactoryList(count int, t *testing.T) []*domain.Item {
	var ItemDomains []*domain.Item
	repo := postgres.NewItemRepository()

	for i := 0; i < count; i++ {
		Item := &Item{}
		err := faker.FakeData(Item)
		if err != nil {
			panic(err)
		}

		ItemDomain, errRepo := repo.Create(Item.ToDomain())
		if errRepo != nil {
			panic(err)
		}
		ItemDomains = append(ItemDomains, ItemDomain)
	}

	t.Cleanup(func() {
		CleanItem()
	})

	return ItemDomains
}

func CleanItem() {
	postgres.NewPostgresBase().DB.Exec("DELETE FROM items")
	postgres.NewPostgresBase().DB.Exec("ALTER SEQUENCE items_id_seq RESTART WITH 1")
}
