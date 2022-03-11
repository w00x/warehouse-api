package domain

import (
	"time"
)

type Price struct {
	Id			uint
	MarketId	uint
	Market 		*Market
	ItemId		uint
	Item        *Item
	Price 		int
	Date 		time.Time
}

func NewPrice(id uint, market *Market, item *Item, price int, date time.Time, ItemId uint, MarketId uint) *Price {
	return &Price{Id: id, Market: market, Item: item, Price: price, Date: date, MarketId: MarketId, ItemId: ItemId}
}
