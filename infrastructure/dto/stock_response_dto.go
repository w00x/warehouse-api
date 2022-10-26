package dto

import (
	"warehouse/domain"
	"warehouse/shared"
)

type StockResponseDto struct {
	Id             string          `json:"id" uri:"id"`
	Item           *ItemDto        `json:"item"`
	Rack           *RackDto        `json:"rack"`
	Quantity       int             `json:"quantity"`
	OperationDate  shared.DateTime `json:"operation_date"`
	ExpirationDate shared.DateTime `json:"expiration_date"`
	Comment        string          `json:"comment"`
}

func NewStockResponseDto(id string, stock *ItemDto, rack *RackDto, quantity int,
	operationDate shared.DateTime, expirationDate shared.DateTime, comment string) *StockResponseDto {
	return &StockResponseDto{Id: id, Item: stock, Rack: rack, Quantity: quantity,
		OperationDate: operationDate, ExpirationDate: expirationDate, Comment: comment}
}

func NewStockResponseDtoFromDomain(stock *domain.Stock) *StockResponseDto {
	return NewStockResponseDto(stock.Id(), NewItemDtoFromDomain(stock.Item),
		NewRackDtoFromDomain(stock.Rack), stock.Quantity,
		stock.OperationDate, stock.ExpirationDate, stock.Comment)
}

func NewStockListResponseDtoFromDomains(stocks *[]domain.Stock) []*StockResponseDto {
	var stockDtos []*StockResponseDto
	for _, stock := range *stocks {
		stockDtos = append(stockDtos,
			NewStockResponseDtoFromDomain(&stock))
	}
	return stockDtos
}

func (s StockResponseDto) ToDomain() *domain.Stock {
	return domain.NewStock(s.Id, s.Item.ToDomain(), s.Rack.ToDomain(), s.Quantity,
		s.OperationDate, s.ExpirationDate, s.Comment, s.Item.Id, s.Rack.Id)
}
