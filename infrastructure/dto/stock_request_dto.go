package dto

import (
	"warehouse/domain"
	"warehouse/shared"
)

type StockRequestDto struct {
	Id             string          `json:"id" uri:"id"`
	ItemId         string          `json:"item_id"`
	RackId         string          `json:"rack_id"`
	Quantity       int             `json:"quantity"`
	OperationDate  shared.DateTime `json:"operation_date"`
	ExpirationDate shared.DateTime `json:"expiration_date"`
	Comment        string          `json:"comment"`
}

func (s StockRequestDto) ToDomain() *domain.Stock {
	return domain.NewStock(s.Id, nil, nil, s.Quantity,
		s.OperationDate, s.ExpirationDate, s.Comment, s.ItemId, s.RackId)
}
