package factories

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"testing"
	"warehouse/domain"
	"warehouse/infraestructure/repository/postgres"
)

type Market struct {
	Name 					string
}

func (i Market) ToDomain() *domain.Market {
	return domain.NewMarket(0, i.Name)
}

func FromMarketDomainToFactory(market *domain.Market) *Market {
	return &Market{Name: market.Name}
}

func NewMarketFactory(t *testing.T) *domain.Market {
	Market := &Market{}
	err := faker.FakeData(Market)
	if err != nil {
		fmt.Println(err)
	}

	repo := postgres.NewMarketRepository()
	MarketDomain, errRepo := repo.Create(Market.ToDomain())
	if errRepo != nil {
		panic(err)
	}

	t.Cleanup(func() {
		CleanMarket()
	})

	return MarketDomain
}

func NewMarketObjectFactory() map[string]interface{} {
	market := &Market{}
	err := faker.FakeData(market)
	if err != nil {
		fmt.Println(err)
	}

	marketMarshal := map[string]interface{}{
		"name": market.Name,
	}

	return marketMarshal
}

func NewMarketFactoryList(count int, t *testing.T) []*domain.Market {
	var MarketDomains []*domain.Market
	repo := postgres.NewMarketRepository()

	for i := 0; i < count; i++ {
		Market := &Market{}
		err := faker.FakeData(Market)
		if err != nil {
			panic(err)
		}

		MarketDomain, errRepo := repo.Create(Market.ToDomain())
		if errRepo != nil {
			panic(err)
		}
		MarketDomains = append(MarketDomains, MarketDomain)
	}

	t.Cleanup(func() {
		CleanMarket()
	})

	return MarketDomains
}

func CleanMarket() {
	postgres.NewPostgresBase().DB.Exec("DELETE FROM markets")
	postgres.NewPostgresBase().DB.Exec("ALTER SEQUENCE markets_id_seq RESTART WITH 1")
}
