package factories

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"testing"
	"warehouse/domain"
	"warehouse/infraestructure/repository/postgres"
	"warehouse/shared"
)

type Stock struct {
	Id				uint
	ItemId			uint
	Item 			*Item
	RackId			uint
	Rack 			*Rack
	Quantity        int
	OperationDate  	shared.DateTime
	ExpirationDate 	shared.DateTime
}

func (i Stock) ToDomain() *domain.Stock {
	return domain.NewStock(0, i.Item.ToDomain(), i.Rack.ToDomain(), i.Quantity, i.OperationDate, i.ExpirationDate, i.ItemId, i.RackId)
}

func NewStockFactory(t *testing.T) *domain.Stock {
	stock := &Stock{}
	err := faker.FakeData(stock)
	if err != nil {
		fmt.Println(err)
	}

	rack := NewRackFactory(t)
	item := NewItemFactory(t)
	stock.Rack = FromRackDomainToFactory(rack)
	stock.Item = FromItemDomainToFactory(item)

	repo := postgres.NewStockRepository()
	StockDomain, errRepo := repo.Create(stock.ToDomain())
	if errRepo != nil {
		panic(err)
	}

	t.Cleanup(func() {
		CleanStock()
	})

	return StockDomain
}

func NewStockObjectFactory() map[string]interface{} {
	stock := &Stock{}
	err := faker.FakeData(stock)
	if err != nil {
		fmt.Println(err)
	}

	stockMarshal := map[string]interface{}{
		"quantity": stock.Quantity,
	}

	return stockMarshal
}

func NewStockObjectForCreateFactory(t *testing.T) map[string]interface{} {
	rack := NewRackFactory(t)
	item := NewItemFactory(t)


	stockMarshal := NewStockObjectFactory()
	stockMarshal["rack_id"] = rack.Id
	stockMarshal["item_id"] = item.Id

	return stockMarshal
}

func NewStockFactoryList(count int, t *testing.T) []*domain.Stock {
	var StockDomains []*domain.Stock
	repo := postgres.NewStockRepository()

	for i := 0; i < count; i++ {
		Stock := &Stock{}
		err := faker.FakeData(Stock)
		if err != nil {
			panic(err)
		}

		StockDomain, errRepo := repo.Create(Stock.ToDomain())
		if errRepo != nil {
			panic(err)
		}
		StockDomains = append(StockDomains, StockDomain)
	}

	t.Cleanup(func() {
		CleanStock()
	})

	return StockDomains
}

func CleanStock() {
	postgres.NewPostgresBase().DB.Exec("DELETE FROM stocks")
	postgres.NewPostgresBase().DB.Exec("ALTER SEQUENCE stocks_id_seq RESTART WITH 1")
}
