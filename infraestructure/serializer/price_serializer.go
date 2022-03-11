package serializer

import (
	"time"
	"warehouse/domain"
	"warehouse/shared"
)

type PriceSerializer struct {
	Id			uint 				`json:"id" uri:"id"`
	Market 		*MarketSerializer	`json:"market"`
	Item        *ItemSerializer		`json:"item"`
	MarketId 	uint				`json:"market_id"`
	ItemId      uint				`json:"item_id"`
	Price 		int					`json:"price"`
	Date 		shared.DateTime		`json:"date"`
}

func NewPriceSerializer(id uint, market *MarketSerializer, item *ItemSerializer, price int,
	date shared.DateTime, ItemId uint, MarketId uint) *PriceSerializer {
	return &PriceSerializer{Id: id, Market: market, Item: item, Price: price, Date: date, ItemId: ItemId, MarketId: MarketId}
}

func NewPriceSerializerFromDomain(price *domain.Price) *PriceSerializer {
	return NewPriceSerializer(price.Id, NewMarketSerializerFromDomain(price.Market),
		NewItemSerializerFromDomain(price.Item), price.Price, shared.DateTime(price.Date), price.ItemId, price.MarketId)
}

func NewPriceListSerializerFromDomains(prices *[]domain.Price) []*PriceSerializer {
	var priceSerializers []*PriceSerializer
	for _, price := range *prices {
		priceSerializers = append(priceSerializers,
			NewPriceSerializerFromDomain(&price))
	}
	return priceSerializers
}

func (ps PriceSerializer) ToDomain() *domain.Price {
	var market *domain.Market
	if ps.Market == nil {
		market = nil
	} else {
		market = ps.Market.ToDomain()
	}

	var item *domain.Item
	if ps.Item == nil {
		item = nil
	} else {
		item = ps.Item.ToDomain()
	}
	return domain.NewPrice(ps.Id, market, item, ps.Price,
		time.Time(ps.Date), ps.ItemId, ps.MarketId)
}
