package dto

import (
	"warehouse/domain"
	"warehouse/shared"
)

type PriceRequesDto struct {
	Id			uint 				`json:"id" uri:"id"`
	MarketId 	uint				`json:"market_id"`
	ItemId      uint				`json:"item_id"`
	Price 		float64					`json:"price"`
	Date 		shared.DateTime		`json:"date"`
}

func (ps PriceRequesDto) ToDomain() *domain.Price {
	return domain.NewPrice(ps.Id, nil, nil, ps.Price, ps.Date, ps.ItemId, ps.MarketId)
}
