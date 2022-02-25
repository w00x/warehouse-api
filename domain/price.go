package domain

import (
	"time"
)

type Price struct {
	Id			string
	Market 		*Market
	Item        *Item
	Price 		int
	Date 		time.Time
}

func NewPrice(id string, market *Market, item *Item, price int, date time.Time) *Price {
	return &Price{Id: id, Market: market, Item: item, Price: price, Date: date}
}
