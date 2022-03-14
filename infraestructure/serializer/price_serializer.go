package serializer

import (
	"time"
	"warehouse/domain"
	"warehouse/shared"
)

type PriceSerializer struct {
	Id			uint 				`json:"id" uri:"id"`
	MarketId 	uint				`json:"market_id"`
	ItemId      uint				`json:"item_id"`
	Price 		int					`json:"price"`
	Date 		shared.DateTime		`json:"date"`
}

func (ps PriceSerializer) ToDomain() *domain.Price {
	return domain.NewPrice(ps.Id, nil, nil, ps.Price, time.Time(ps.Date), ps.ItemId, ps.MarketId)
}
