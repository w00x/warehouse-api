package serializer

import (
	"time"
	"warehouse/domain"
	"warehouse/shared"
)

type StockSerializer struct {
	Id				string
	Item 			*ItemSerializer
	Rack 			*RackSerializer
	Quantity        int
	OperationDate  	shared.DateTime
	ExpirationDate 	shared.DateTime
}

func NewStockSerializer(id string, item *ItemSerializer, rack *RackSerializer, quantity int,
	operationDate shared.DateTime, expirationDate shared.DateTime) *StockSerializer {
	return &StockSerializer{Id: id, Item: item, Rack: rack, Quantity: quantity,
		OperationDate: operationDate, ExpirationDate: expirationDate}
}

func NewStockSerializerFromDomain(stock *domain.Stock) *StockSerializer {
	return NewStockSerializer(stock.Id, NewItemSerializerFromDomain(stock.Item),
		NewRackSerializerFromDomain(stock.Rack), stock.Quantity,
		shared.DateTime(stock.OperationDate), shared.DateTime(stock.ExpirationDate))
}

func (s StockSerializer) ToDomain() *domain.Stock {
	return domain.NewStock(s.Id, s.Item.ToDomain(), s.Rack.ToDomain(), s.Quantity,
		time.Time(s.OperationDate), time.Time(s.ExpirationDate))
}
