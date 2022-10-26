package dto

import (
	"warehouse/domain"
	"warehouse/shared"
)

type PriceRequesDto struct {
	Id       string          `json:"id" uri:"id"`
	MarketId string          `json:"market_id"`
	ItemId   string          `json:"item_id"`
	Price    float64         `json:"price"`
	Date     shared.DateTime `json:"date"`
}

func (ps PriceRequesDto) ToDomain() *domain.Price {
	return domain.NewPrice(ps.Id, nil, nil, ps.Price, ps.Date, ps.ItemId, ps.MarketId)
}
