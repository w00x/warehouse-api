package dto

import (
	"warehouse/domain"
	"warehouse/shared"
)

type StockRequestDto struct {
	Id				uint 				`json:"id" uri:"id"`
	ItemId 			uint				`json:"item_id"`
	RackId 			uint				`json:"rack_id"`
	Quantity        int					`json:"quantity"`
	OperationDate  	shared.DateTime		`json:"operation_date"`
	ExpirationDate 	shared.DateTime		`json:"expiration_date"`
}

func (s StockRequestDto) ToDomain() *domain.Stock {
	return domain.NewStock(s.Id, nil, nil, s.Quantity,
		s.OperationDate, s.ExpirationDate, s.ItemId, s.RackId)
}
