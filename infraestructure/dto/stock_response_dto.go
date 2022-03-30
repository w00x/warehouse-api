package dto

import (
	"warehouse/domain"
	"warehouse/shared"
)

type StockResponseDto struct {
	Id				uint 				`json:"id" uri:"id"`
	Item 			*ItemDto		`json:"item"`
	Rack 			*RackDto		`json:"rack"`
	Quantity        int					`json:"quantity"`
	OperationDate  	shared.DateTime		`json:"operation_date"`
	ExpirationDate 	shared.DateTime		`json:"expiration_date"`
}

func NewStockResponseDto(id uint, stock *ItemDto, rack *RackDto, quantity int,
	operationDate shared.DateTime, expirationDate shared.DateTime) *StockResponseDto {
	return &StockResponseDto{Id: id, Item: stock, Rack: rack, Quantity: quantity,
		OperationDate: operationDate, ExpirationDate: expirationDate}
}

func NewStockResponseDtoFromDomain(stock *domain.Stock) *StockResponseDto {
	return NewStockResponseDto(stock.Id, NewItemDtoFromDomain(stock.Item),
		NewRackDtoFromDomain(stock.Rack), stock.Quantity,
		shared.DateTime(stock.OperationDate), shared.DateTime(stock.ExpirationDate))
}

func NewStockListResponseDtoFromDomains(stocks *[]domain.Stock) []*StockResponseDto {
	var stockDtos []*StockResponseDto
	for _, stock := range *stocks {
		stockDtos = append(stockDtos,
			NewStockResponseDtoFromDomain(&stock))
	}
	return stockDtos
}