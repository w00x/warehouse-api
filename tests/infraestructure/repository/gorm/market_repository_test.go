package gorm

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"testing"
	"warehouse/infrastructure/repository/gorm"
	"warehouse/tests/factories"
)

func TestMarketRepository_All(t *testing.T) {
	sizeOfMarkets := 5
	markets := factories.NewMarketFactoryList(sizeOfMarkets)
	marketRepo := gorm.NewMarketRepository()
	allMarkets, err := marketRepo.All()
	assert.Nil(t, err)

	var marketsIds []string
	for _, market := range *allMarkets {
		marketsIds = append(marketsIds, market.Id())
	}

	assert.Contains(t, marketsIds, markets[0].Id())
	assert.Contains(t, marketsIds, markets[1].Id())
	assert.Contains(t, marketsIds, markets[2].Id())
}

func TestMarketRepository_Create(t *testing.T) {
	marketRepo := gorm.NewMarketRepository()
	marketData := factories.NewMarketDomainFactory()

	market, err := marketRepo.Create(marketData)
	assert.Nil(t, err)

	assert.Equal(t, market.Name, marketData.Name)
	assert.NotNil(t, market.Id())
}

func TestMarketRepository_Delete(t *testing.T) {
	market := factories.NewMarketFactory()
	marketId := market.Id()
	assert.NotNil(t, marketId)

	marketRepo := gorm.NewMarketRepository()
	marketFounded, err := marketRepo.Find(marketId)
	assert.Nil(t, err)
	assert.NotNil(t, marketFounded)
	assert.Equal(t, marketId, marketFounded.Id())

	assert.Nil(t, marketRepo.Delete(market))
	marketFounded, err = marketRepo.Find(marketId)
	assert.NotNil(t, err)
	assert.Nil(t, marketFounded)
}

func TestMarketRepository_Find(t *testing.T) {
	market := factories.NewMarketFactory()
	marketId := market.Id()
	assert.NotNil(t, marketId)

	marketRepo := gorm.NewMarketRepository()
	marketFounded, err := marketRepo.Find(marketId)
	assert.Nil(t, err)
	assert.NotNil(t, marketFounded)
	assert.Equal(t, marketId, marketFounded.Id())
}

func TestMarketRepository_Update(t *testing.T) {
	market := factories.NewMarketFactory()
	newName := gofakeit.LoremIpsumWord()
	marketRepo := gorm.NewMarketRepository()
	market.Name = newName
	marketUpdated, err := marketRepo.Update(market)

	assert.Nil(t, err)
	assert.NotNil(t, marketUpdated)

	marketFounded, err := marketRepo.Find(marketUpdated.Id())
	assert.Nil(t, err)
	assert.NotNil(t, marketFounded)
	assert.Equal(t, newName, marketFounded.Name)
}

func TestNewMarketRepository(t *testing.T) {
	repo := gorm.NewMarketRepository()
	assert.NotNil(t, repo)
	assert.IsType(t, gorm.MarketRepository{}, *repo)
}
