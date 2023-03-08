package factories

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"testing"
	"warehouse/domain"
	"warehouse/infrastructure/repository/gorm"
	"warehouse/shared"
)

type Price struct {
	Market   *Market
	Item     *Item
	Price    float64 `fake:"{float32}"`
	Date     shared.DateTime
	ItemId   string
	MarketId string
}

func (i Price) ToDomain() *domain.Price {
	return domain.NewPrice("", i.Market.ToDomain(), i.Item.ToDomain(), i.Price, i.Date, i.ItemId, i.MarketId)
}

func NewPriceFactory() *domain.Price {
	price := &Price{}
	err := gofakeit.Struct(price)
	if err != nil {
		fmt.Println(err)
	}

	market := NewMarketFactory()
	item := NewItemFactory()
	price.Market = FromMarketDomainToFactory(market)
	price.Item = FromItemDomainToFactory(item)

	repo := gorm.NewPriceRepository()
	PriceDomain, errRepo := repo.Create(price.ToDomain())
	if errRepo != nil {
		panic(err)
	}

	return PriceDomain
}

func NewPriceDomainFactory() *domain.Price {
	price := &Price{}
	err := gofakeit.Struct(price)
	if err != nil {
		fmt.Println(err)
	}

	market := NewMarketFactory()
	item := NewItemFactory()
	price.Market = FromMarketDomainToFactory(market)
	price.Item = FromItemDomainToFactory(item)

	return price.ToDomain()
}

func NewPriceObjectFactory() map[string]interface{} {
	price := &Price{}
	err := gofakeit.Struct(price)
	if err != nil {
		fmt.Println(err)
	}

	priceMarshal := map[string]interface{}{
		"price": price.Price,
		"date":  price.Date.Format("2006-01-02 15:04:05"),
	}

	return priceMarshal
}

func NewPriceObjectForCreateFactory(t *testing.T) map[string]interface{} {
	market := NewMarketFactory()
	item := NewItemFactory()

	priceMarshal := NewPriceObjectFactory()
	priceMarshal["market_id"] = market.Id()
	priceMarshal["item_id"] = item.Id()

	return priceMarshal
}

func NewPriceFactoryList(count int) []*domain.Price {
	var PriceDomains []*domain.Price
	repo := gorm.NewPriceRepository()

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

	return PriceDomains
}
