package dto

import (
	"warehouse/shared"
)

type StockRequestDto struct {
	Id             string          `json:"id" uri:"id"`
	ItemCode       string          `json:"item_code"`
	RackCode       string          `json:"rack_code"`
	Quantity       int             `json:"quantity"`
	OperationDate  shared.DateTime `json:"operation_date"`
	ExpirationDate shared.DateTime `json:"expiration_date"`
	Comment        string          `json:"comment"`
}
