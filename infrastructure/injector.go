package infrastructure

import (
	"warehouse/application"
	"warehouse/infrastructure/controller/v1"
	"warehouse/infrastructure/factory/repository"
)

func InitializeInventoryController(factoryAdapter string) *v1.InventoryController {
	repository, err := repository.InventoryFactory(factoryAdapter)
	if err != nil {
		panic(err)
	}
	application := application.NewInventoryApplication(repository)
	return v1.NewInventoryController(application)
}

func InitializeItemController(factoryAdapter string) *v1.ItemController {
	repository, err := repository.ItemFactory(factoryAdapter)
	if err != nil {
		panic(err)
	}
	application := application.NewItemApplication(repository)
	return v1.NewItemController(application)
}

func InitializeMarketController(factoryAdapter string) *v1.MarketController {
	repository, err := repository.MarketFactory(factoryAdapter)
	if err != nil {
		panic(err)
	}
	application := application.NewMarketApplication(repository)
	return v1.NewMarketController(application)
}

func InitializePriceController(factoryAdapter string) *v1.PriceController {
	repository, err := repository.PriceFactory(factoryAdapter)
	if err != nil {
		panic(err)
	}
	application := application.NewPriceApplication(repository)
	return v1.NewPriceController(application)
}

func InitializeRackController(factoryAdapter string) *v1.RackController {
	repository, err := repository.RackFactory(factoryAdapter)
	if err != nil {
		panic(err)
	}
	application := application.NewRackApplication(repository)
	return v1.NewRackController(application)
}

func InitializeStockController(factoryAdapter string) *v1.StockController {
	stockRepository, err := repository.StockFactory(factoryAdapter)
	itemRepository, err := repository.ItemFactory(factoryAdapter)
	rackRepository, err := repository.RackFactory(factoryAdapter)
	if err != nil {
		panic(err)
	}
	application := application.NewStockApplication(stockRepository, itemRepository, rackRepository)
	return v1.NewStockController(application)
}
