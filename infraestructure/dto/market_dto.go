package dto

import "warehouse/domain"

type MarketDto struct {
	Id		uint 	`json:"id" uri:"id"`
	Name 	string	`json:"name"`
}

func NewMarketDto(id uint, name string) *MarketDto {
	return &MarketDto{Id: id, Name: name}
}

func NewMarketDtoFromDomain(market *domain.Market) *MarketDto {
	if market == nil {
		return nil
	}
	return NewMarketDto(market.Id, market.Name)
}

func NewMarketListDtoFromDomains(markets *[]domain.Market) []*MarketDto {
	var marketDtos []*MarketDto
	for _, market := range *markets {
		marketDtos = append(marketDtos,
			NewMarketDtoFromDomain(&market))
	}
	return marketDtos
}

func (m MarketDto) ToDomain() *domain.Market {
	return domain.NewMarket(m.Id, m.Name)
}