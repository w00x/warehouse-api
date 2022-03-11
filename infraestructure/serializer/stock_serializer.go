package serializer

import (
	"time"
	"warehouse/domain"
	"warehouse/shared"
)

type StockSerializer struct {
	Id				uint 				`json:"id" uri:"id"`
	Item 			*ItemSerializer		`json:"item"`
	Rack 			*RackSerializer		`json:"rack"`
	ItemId 			uint				`json:"item_id"`
	RackId 			uint				`json:"rack_id"`
	Quantity        int					`json:"quantity"`
	OperationDate  	shared.DateTime		`json:"operation_date"`
	ExpirationDate 	shared.DateTime		`json:"expiration_date"`
}

func NewStockSerializer(id uint, stock *ItemSerializer, rack *RackSerializer, quantity int,
	operationDate shared.DateTime, expirationDate shared.DateTime, ItemId uint, RakId uint) *StockSerializer {
	return &StockSerializer{Id: id, Item: stock, Rack: rack, Quantity: quantity,
		OperationDate: operationDate, ExpirationDate: expirationDate, ItemId: ItemId, RackId: RakId}
}

func NewStockSerializerFromDomain(stock *domain.Stock) *StockSerializer {
	return NewStockSerializer(stock.Id, NewItemSerializerFromDomain(stock.Item),
		NewRackSerializerFromDomain(stock.Rack), stock.Quantity,
		shared.DateTime(stock.OperationDate), shared.DateTime(stock.ExpirationDate), stock.ItemId, stock.RackId)
}

func NewStockListSerializerFromDomains(stocks *[]domain.Stock) []*StockSerializer {
	var stockSerializers []*StockSerializer
	for _, stock := range *stocks {
		stockSerializers = append(stockSerializers,
			NewStockSerializerFromDomain(&stock))
	}
	return stockSerializers
}

func (s StockSerializer) ToDomain() *domain.Stock {
	var rack *domain.Rack
	if s.Rack == nil {
		rack = nil
	} else {
		rack = s.Rack.ToDomain()
	}

	var item *domain.Item
	if s.Item == nil {
		item = nil
	} else {
		item = s.Item.ToDomain()
	}
	return domain.NewStock(s.Id, item, rack, s.Quantity,
		time.Time(s.OperationDate), time.Time(s.ExpirationDate), s.ItemId, s.RackId)
}
