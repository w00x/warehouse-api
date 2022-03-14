package serializer

import (
	"warehouse/domain"
	"warehouse/shared"
)

type StockResponseSerializer struct {
	Id				uint 				`json:"id" uri:"id"`
	Item 			*ItemSerializer		`json:"item"`
	Rack 			*RackSerializer		`json:"rack"`
	Quantity        int					`json:"quantity"`
	OperationDate  	shared.DateTime		`json:"operation_date"`
	ExpirationDate 	shared.DateTime		`json:"expiration_date"`
}

func NewStockResponseSerializer(id uint, stock *ItemSerializer, rack *RackSerializer, quantity int,
	operationDate shared.DateTime, expirationDate shared.DateTime) *StockResponseSerializer {
	return &StockResponseSerializer{Id: id, Item: stock, Rack: rack, Quantity: quantity,
		OperationDate: operationDate, ExpirationDate: expirationDate}
}

func NewStockResponseSerializerFromDomain(stock *domain.Stock) *StockResponseSerializer {
	return NewStockResponseSerializer(stock.Id, NewItemSerializerFromDomain(stock.Item),
		NewRackSerializerFromDomain(stock.Rack), stock.Quantity,
		shared.DateTime(stock.OperationDate), shared.DateTime(stock.ExpirationDate))
}

func NewStockListResponseSerializerFromDomains(stocks *[]domain.Stock) []*StockResponseSerializer {
	var stockSerializers []*StockResponseSerializer
	for _, stock := range *stocks {
		stockSerializers = append(stockSerializers,
			NewStockResponseSerializerFromDomain(&stock))
	}
	return stockSerializers
}