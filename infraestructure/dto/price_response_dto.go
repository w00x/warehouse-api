package dto

import (
	"warehouse/domain"
	"warehouse/shared"
)

type PriceResponseDto struct {
	Id			uint 				`json:"id" uri:"id"`
	Market 		*MarketDto	`json:"market"`
	Item        *ItemDto		`json:"item"`
	Price 		float64				`json:"price"`
	Date 		shared.DateTime		`json:"date"`
}

func NewPriceResponseDto(id uint, market *MarketDto, item *ItemDto, price float64,
	date shared.DateTime) *PriceResponseDto {
	return &PriceResponseDto{Id: id, Market: market, Item: item, Price: price, Date: date}
}

func NewPriceResponseDtoFromDomain(price *domain.Price) *PriceResponseDto {
	return NewPriceResponseDto(price.Id, NewMarketDtoFromDomain(price.Market),
		NewItemDtoFromDomain(price.Item), price.Price, shared.DateTime(price.Date))
}

func NewPriceResponseListDtoFromDomains(prices *[]domain.Price) []*PriceResponseDto {
	var priceDtos []*PriceResponseDto
	for _, price := range *prices {
		priceDtos = append(priceDtos,
			NewPriceResponseDtoFromDomain(&price))
	}
	return priceDtos
}