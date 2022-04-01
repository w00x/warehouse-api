package factories

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"testing"
	"time"
	"warehouse/domain"
	"warehouse/infraestructure/repository/postgres"
	"warehouse/shared"
)

type Price struct {
	Market *Market
	Item *Item
	Price float64
	Date shared.DateTime
	ItemId uint
	MarketId uint
}

func (i Price) ToDomain() *domain.Price {
	return domain.NewPrice(0, i.Market.ToDomain(), i.Item.ToDomain(), i.Price, i.Date, i.ItemId, i.MarketId)
}

func NewPriceFactory(t *testing.T) *domain.Price {
	price := &Price{}
	err := gofakeit.Struct(price)
	if err != nil {
		fmt.Println(err)
	}

	market := NewMarketFactory(t)
	item := NewItemFactory(t)
	price.Market = FromMarketDomainToFactory(market)
	price.Item = FromItemDomainToFactory(item)

	repo := postgres.NewPriceRepository()
	PriceDomain, errRepo := repo.Create(price.ToDomain())
	if errRepo != nil {
		panic(err)
	}

	t.Cleanup(func() {
		CleanPrice()
	})

	return PriceDomain
}

func NewPriceObjectFactory() map[string]interface{} {
	price := &Price{}
	err := gofakeit.Struct(price)
	if err != nil {
		fmt.Println(err)
	}

	priceMarshal := map[string]interface{}{
		"price": price.Price,
		"date": time.Time(price.Date).Format("2006-01-02 15:04:05"),
	}

	return priceMarshal
}

func NewPriceObjectForCreateFactory(t *testing.T) map[string]interface{} {
	market := NewMarketFactory(t)
	item := NewItemFactory(t)


	priceMarshal := NewPriceObjectFactory()
	priceMarshal["market_id"] = market.Id
	priceMarshal["item_id"] = item.Id

	return priceMarshal
}

func NewPriceFactoryList(count int, t *testing.T) []*domain.Price {
	var PriceDomains []*domain.Price
	repo := postgres.NewPriceRepository()

	for i := 0; i < count; i++ {
		Price := &Price{}
		err := gofakeit.Struct(Price)
		if err != nil {
			panic(err)
		}

		PriceDomain, errRepo := repo.Create(Price.ToDomain())
		if errRepo != nil {
			panic(err)
		}
		PriceDomains = append(PriceDomains, PriceDomain)
	}

	t.Cleanup(func() {
		CleanPrice()
	})

	return PriceDomains
}

func CleanPrice() {
	postgres.NewPostgresBase().DB.Exec("DELETE FROM prices")
	postgres.NewPostgresBase().DB.Exec("ALTER SEQUENCE prices_id_seq RESTART WITH 1")
}
