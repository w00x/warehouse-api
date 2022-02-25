package serializer

import "warehouse/domain"

type MarketSerializer struct {
	Id		string
	Name 	string
}

func NewMarketSerializer(id string, name string) *MarketSerializer {
	return &MarketSerializer{Id: id, Name: name}
}

func NewMarketSerializerFromDomain(market *domain.Market) *MarketSerializer {
	return NewMarketSerializer(market.Id, market.Name)
}

func (m MarketSerializer) ToDomain() *domain.Market {
	return domain.NewMarket(m.Id, m.Name)
}