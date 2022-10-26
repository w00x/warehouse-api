package models

import (
	"gorm.io/gorm"
	"warehouse/shared"
)

type Price struct {
	gorm.Model
	ID       string `gorm:"type:uuid;primaryKey;column:id"`
	MarketId string
	Market   *Market
	ItemId   string
	Item     *Item
	Price    float64
	Date     shared.DateTime
}

func NewPrice(id string, market *Market, item *Item, price float64, date shared.DateTime, ItemId string, MarketId string) *Price {
	return &Price{ID: id, Market: market, Item: item, Price: price, Date: date, ItemId: ItemId, MarketId: MarketId}
}
