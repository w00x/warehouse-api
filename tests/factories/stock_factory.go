package factories

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"warehouse/domain"
	"warehouse/infrastructure/repository/gorm"
	"warehouse/shared"
)

type Stock struct {
	Id             string
	ItemId         string
	Item           *Item
	RackId         string
	Rack           *Rack
	Quantity       int `fake:"{number:1,10}"`
	OperationDate  shared.DateTime
	ExpirationDate shared.DateTime
	Comment        string
}

func (i Stock) ToDomain() *domain.Stock {
	return domain.NewStock("", i.Item.ToDomain(), i.Rack.ToDomain(), i.Quantity,
		i.OperationDate, i.ExpirationDate, i.Comment, i.ItemId, i.RackId)
}

func NewStockFactory() *domain.Stock {
	stock := &Stock{}
	err := gofakeit.Struct(stock)
	if err != nil {
		fmt.Println(err)
	}

	rack := NewRackFactory()
	item := NewItemFactory()
	stock.Rack = FromRackDomainToFactory(rack)
	stock.Item = FromItemDomainToFactory(item)

	repo := gorm.NewStockRepository()
	StockDomain, errRepo := repo.Create(stock.ToDomain())
	if errRepo != nil {
		panic(err)
	}

	return StockDomain
}

func NewStockDomainFactory() *domain.Stock {
	stock := &Stock{}
	err := gofakeit.Struct(stock)
	if err != nil {
		fmt.Println(err)
	}

	rack := NewRackFactory()
	item := NewItemFactory()
	stock.Rack = FromRackDomainToFactory(rack)
	stock.Item = FromItemDomainToFactory(item)

	return stock.ToDomain()
}

func NewStockObjectFactory() map[string]interface{} {
	stock := &Stock{}
	err := gofakeit.Struct(stock)
	if err != nil {
		fmt.Println(err)
	}

	stockMarshal := map[string]interface{}{
		"quantity":        stock.Quantity,
		"operation_date":  stock.OperationDate.Format("2006-01-02 15:04:05"),
		"expiration_date": stock.ExpirationDate.Format("2006-01-02 15:04:05"),
		"comment":         stock.Comment,
	}

	return stockMarshal
}

func NewStockObjectForCreateFactory() map[string]interface{} {
	rack := NewRackFactory()
	item := NewItemFactory()

	stockMarshal := NewStockObjectFactory()
	stockMarshal["rack_code"] = rack.Code
	stockMarshal["item_code"] = item.Code

	return stockMarshal
}

func NewStockFactoryList(count int) []*domain.Stock {
	var StockDomains []*domain.Stock
	repo := gorm.NewStockRepository()

	for i := 0; i < count; i++ {
		stock := &Stock{}
		err := gofakeit.Struct(stock)
		if err != nil {
			panic(err)
		}
		stock.OperationDate = shared.TimeToDateTime(gofakeit.Date())
		StockDomain, errRepo := repo.Create(stock.ToDomain())
		if errRepo != nil {
			panic(err)
		}
		StockDomains = append(StockDomains, StockDomain)
	}

	return StockDomains
}
