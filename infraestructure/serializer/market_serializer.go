package serializer

import "warehouse/domain"

type MarketSerializer struct {
	Id		uint 	`json:"id" uri:"id"`
	Name 	string	`json:"name"`
}

func NewMarketSerializer(id uint, name string) *MarketSerializer {
	return &MarketSerializer{Id: id, Name: name}
}

func NewMarketSerializerFromDomain(market *domain.Market) *MarketSerializer {
	if market == nil {
		return nil
	}
	return NewMarketSerializer(market.Id, market.Name)
}

func NewMarketListSerializerFromDomains(markets *[]domain.Market) []*MarketSerializer {
	var marketSerializers []*MarketSerializer
	for _, market := range *markets {
		marketSerializers = append(marketSerializers,
			NewMarketSerializerFromDomain(&market))
	}
	return marketSerializers
}

func (m MarketSerializer) ToDomain() *domain.Market {
	return domain.NewMarket(m.Id, m.Name)
}