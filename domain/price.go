package domain

import (
	"github.com/google/uuid"
	"warehouse/shared"
)

type Price struct {
	id       string
	MarketId string
	Market   *Market
	ItemId   string
	Item     *Item
	Price    float64
	Date     shared.DateTime
}

func NewPrice(id string, market *Market, item *Item, price float64, date shared.DateTime, ItemId string, MarketId string) *Price {
	if id == "" {
		id = uuid.New().String()
	}

	return &Price{id: id, Market: market, Item: item, Price: price, Date: date, MarketId: MarketId,
		ItemId: ItemId}
}

func (i Price) Id() string {
	return i.id
}
