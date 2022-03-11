package models

import (
	"gorm.io/gorm"
	"time"
	"warehouse/domain"
)

type Price struct {
	gorm.Model
	Id			uint		`gorm:"primaryKey;autoIncrement"`
	MarketId	uint
	Market 		*Market
	ItemId		uint
	Item        *Item
	Price 		int
	Date 		time.Time
}

func NewPrice(id uint, market *Market, item *Item, price int, date time.Time, ItemId uint, MarketId uint) *Price {
	return &Price{Id: id, Market: market, Item: item, Price: price, Date: date, ItemId: ItemId, MarketId: MarketId}
}

func FromPriceDomainToModel(i *domain.Price) *Price {
	return NewPrice(i.Id, FromMarketDomainToModel(i.Market), FromItemDomainToModel(i.Item), i.Price, i.Date, i.ItemId, i.MarketId)
}

func (i Price) ToDomain() *domain.Price {
	var market *domain.Market
	var item *domain.Item
	if i.Market != nil {
		market = i.Market.ToDomain()
	}
	if i.Item != nil {
		item = i.Item.ToDomain()
	}
	return domain.NewPrice(i.Id, market, item, i.Price, i.Date, i.ItemId, i.MarketId)
}