package factories

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"warehouse/domain"
	"warehouse/infrastructure/repository/gorm"
)

type Market struct {
	Name string
}

func (i Market) ToDomain() *domain.Market {
	return domain.NewMarket("", i.Name)
}

func FromMarketDomainToFactory(market *domain.Market) *Market {
	return &Market{Name: market.Name}
}

func NewMarketFactory() *domain.Market {
	market := &Market{}
	err := gofakeit.Struct(market)
	if err != nil {
		fmt.Println(err)
	}

	repo := gorm.NewMarketRepository()
	MarketDomain, errRepo := repo.Create(market.ToDomain())
	if errRepo != nil {
		panic(err)
	}

	return MarketDomain
}

func NewMarketDomainFactory() *domain.Market {
	market := &Market{}
	err := gofakeit.Struct(market)
	if err != nil {
		fmt.Println(err)
	}

	return market.ToDomain()
}

func NewMarketObjectFactory() map[string]interface{} {
	market := &Market{}
	err := gofakeit.Struct(market)
	if err != nil {
		fmt.Println(err)
	}

	marketMarshal := map[string]interface{}{
		"name": market.Name,
	}

	return marketMarshal
}

func NewMarketFactoryList(count int) []*domain.Market {
	var MarketDomains []*domain.Market
	repo := gorm.NewMarketRepository()

	for i := 0; i < count; i++ {
		Market := &Market{}
		err := gofakeit.Struct(Market)
		if err != nil {
			panic(err)
		}

		MarketDomain, errRepo := repo.Create(Market.ToDomain())
		if errRepo != nil {
			panic(err)
		}
		MarketDomains = append(MarketDomains, MarketDomain)
	}

	return MarketDomains
}
