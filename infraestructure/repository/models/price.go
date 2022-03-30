package models

import (
	"gorm.io/gorm"
	"warehouse/shared"
)

type Price struct {
	gorm.Model
	Id			uint			`gorm:"primaryKey;autoIncrement"`
	MarketId	uint
	Market 		*Market
	ItemId		uint
	Item        *Item
	Price 		float64
	Date 		shared.DateTime
}

func NewPrice(id uint, market *Market, item *Item, price float64, date shared.DateTime, ItemId uint, MarketId uint) *Price {
	return &Price{Id: id, Market: market, Item: item, Price: price, Date: date, ItemId: ItemId, MarketId: MarketId}
}