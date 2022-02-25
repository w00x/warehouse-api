package serializer

import (
	"time"
	"warehouse/domain"
	"warehouse/shared"
)

type PriceSerializer struct {
	Id			string
	Market 		*MarketSerializer
	Item        *ItemSerializer
	Price 		int
	Date 		shared.DateTime
}

func NewPriceSerializer(id string, market *MarketSerializer, item *ItemSerializer, price int,
	date shared.DateTime) *PriceSerializer {
	return &PriceSerializer{id, market, item, price, date}
}

func NewPriceSerializerFromDomain(price *domain.Price) *PriceSerializer {
	return NewPriceSerializer(price.Id, NewMarketSerializerFromDomain(price.Market),
		NewItemSerializerFromDomain(price.Item), price.Price, shared.DateTime(price.Date))
}

func (ps PriceSerializer) ToDomain() *domain.Price {
	return domain.NewPrice(ps.Id, ps.Market.ToDomain(), ps.Item.ToDomain(), ps.Price,
		time.Time(ps.Date))
}
