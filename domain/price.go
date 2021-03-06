package domain

import (
	"warehouse/shared"
)

type Price struct {
	Id			uint
	MarketId	uint
	Market 		*Market
	ItemId		uint
	Item        *Item
	Price 		float64
	Date 		shared.DateTime
}

func NewPrice(id uint, market *Market, item *Item, price float64, date shared.DateTime, ItemId uint, MarketId uint) *Price {
	return &Price{Id: id, Market: market, Item: item, Price: price, Date: date, MarketId: MarketId, ItemId: ItemId}
}
