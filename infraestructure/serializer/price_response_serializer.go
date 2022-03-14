package serializer

import (
	"warehouse/domain"
	"warehouse/shared"
)

type PriceResponseSerializer struct {
	Id			uint 				`json:"id" uri:"id"`
	Market 		*MarketSerializer	`json:"market"`
	Item        *ItemSerializer		`json:"item"`
	Price 		int					`json:"price"`
	Date 		shared.DateTime		`json:"date"`
}

func NewPriceResponseSerializer(id uint, market *MarketSerializer, item *ItemSerializer, price int,
	date shared.DateTime) *PriceResponseSerializer {
	return &PriceResponseSerializer{Id: id, Market: market, Item: item, Price: price, Date: date}
}

func NewPriceResponseSerializerFromDomain(price *domain.Price) *PriceResponseSerializer {
	return NewPriceResponseSerializer(price.Id, NewMarketSerializerFromDomain(price.Market),
		NewItemSerializerFromDomain(price.Item), price.Price, shared.DateTime(price.Date))
}

func NewPriceResponseListSerializerFromDomains(prices *[]domain.Price) []*PriceResponseSerializer {
	var priceSerializers []*PriceResponseSerializer
	for _, price := range *prices {
		priceSerializers = append(priceSerializers,
			NewPriceResponseSerializerFromDomain(&price))
	}
	return priceSerializers
}